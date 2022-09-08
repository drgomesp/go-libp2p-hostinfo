package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	mr "math/rand"

	"github.com/davecgh/go-spew/spew"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"

	hostinfo "github.com/drgomesp/go-libp2p-hostinfo"
)

func makeBasicHost(listenPort int, randseed int64) (host.Host, error) {
	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mr.New(mr.NewSource(randseed))
	}

	privk, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	check(err)

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(privk),
		libp2p.DisableRelay(),
	}

	opts = append(opts, libp2p.NoSecurity)

	return libp2p.New(opts...)
}

func main() {
	listenPort := flag.Int("p", 0, "wait for incoming connections")
	httpAddr := flag.String("h", ":4000", "HTTP listen address")
	seedF := flag.Int64("seed", 0, "set random seed for id generation")

	flag.Parse()
	ctx := context.Background()

	h, err := makeBasicHost(*listenPort, *seedF)
	check(err)

	svc, err := hostinfo.NewService(
		ctx,
		h,
		hostinfo.WithGrpcGatewayAddr(*httpAddr),
		hostinfo.WithHttpServeMux(runtime.NewServeMux()),
	)

	res, err := svc.Info(ctx, nil)
	spew.Dump(res)

	<-ctx.Done()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
