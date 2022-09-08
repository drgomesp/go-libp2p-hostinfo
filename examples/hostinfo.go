package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	mr "math/rand"

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
	host, err := libp2p.New()
	if err != nil {
		log.Fatal(err)
	}

	// create a mux to handle the grpc gateway requests
	// this can also be combined with other muxes
	mux := runtime.NewServeMux()

	ctx := context.Background()
	svc, err := hostinfo.NewService(
		ctx,
		host,
		hostinfo.WithGrpcGatewayAddr(":4000"),
		hostinfo.WithHttpServeMux(mux),
	)

	go svc.ListenAndServe()

	log.Println("visit: http://localhost:4000/v1/hostinfo")

	<-ctx.Done()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
