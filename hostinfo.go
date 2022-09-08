package hostinfo

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"time"

	libp2pgrpc "github.com/drgomesp/go-libp2p-grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/libp2p/go-libp2p/core/host"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	v1 "github.com/drgomesp/go-libp2p-hostinfo/proto/v1"
)

type Service struct {
	v1.UnimplementedHostInfoServiceServer

	host        host.Host
	mux         *runtime.ServeMux
	gatewayAddr string
}

func WithHttpServeMux(mux *runtime.ServeMux) ServiceOption {
	return func(s *Service) {
		s.mux = mux
	}
}

func NewService(ctx context.Context, host host.Host, opts ...ServiceOption) (*Service, error) {
	svc := &Service{host: host}

	srv, err := libp2pgrpc.NewGrpcServer(
		ctx,
		host,
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    time.Second,
			Timeout: 3 * time.Second,
		}),
	)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(svc)
	}

	v1.RegisterHostInfoServiceServer(srv, svc)
	err = v1.RegisterHostInfoServiceHandlerServer(ctx, svc.mux, svc)

	return svc, err
}

func (s *Service) Info(_ context.Context, _ *v1.InfoRequest) (*v1.InfoResponse, error) {
	peers := make([]string, 0)
	for _, peer := range s.host.Peerstore().Peers() {
		if peer.ShortString() != s.host.ID().ShortString() {
			peers = append(peers, peer.ShortString())
		}
	}

	sort.Strings(peers)

	return &v1.InfoResponse{
		Id: s.host.ID().String(),
		Addresses: func() []string {
			res := make([]string, 0)

			for _, addr := range s.host.Addrs() {
				res = append(res, fmt.Sprintf("%s/p2p/%s", addr.String(), s.host.ID().String()))
			}

			return res
		}(),
		Protocols: s.host.Mux().Protocols(),
		Peers:     peers,
	}, nil
}

func (s *Service) ListenAndServe() error {
	return http.ListenAndServe(s.gatewayAddr, s.mux)
}
