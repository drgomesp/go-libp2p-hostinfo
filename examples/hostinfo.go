package main

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/libp2p/go-libp2p"

	hostinfo "github.com/drgomesp/go-libp2p-hostinfo"
)

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
	if err != nil {
		log.Fatal(err)
	}

	go svc.ListenAndServe()

	log.Println("visit: http://localhost:4000/v1/hostinfo")

	<-ctx.Done()
}
