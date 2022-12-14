# go-libp2p-hostinfo

[![madeby](https://img.shields.io/badge/made%20by-%40drgomesp-blue)](https://github.com/drgomesp/)
[![Go Report Card](https://goreportcard.com/badge/github.com/drgomesp/go-libp2p-hostinfo)](https://goreportcard.com/report/github.com/drgomesp/go-libp2p-hostinfo)
[![build](https://github.com/drgomesp/go-libp2p-hostinfo/actions/workflows/go-test.yml/badge.svg?style=squared)](https://github.com/drgomesp/go-libp2p-grpc/actions)
[![codecov](https://codecov.io/gh/drgomesp/go-libp2p-hostinfo/branch/main/graph/badge.svg?token=BRMFJRJV2X)](https://codecov.io/gh/drgomesp/go-libp2p-hostinfo)

> A lightweight libp2p host service that exposes general information about the host and the network.

## Table of Contents

- [Install](#install)
- [Features](#features)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Install

```bash
go get github.com/drgomesp/go-libp2p-hostinfo
```

## Features

- [x] A gRPC service that exposes general host/network information
- [x] HTTP/OpenAPI gateway support
- [ ] Support for configurable info to expose

## Usage

For a given libp2p host, create a new `hostinfo.Service` and start it:

```go
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
```

You should now be able to access the hostinfo endpoint:

```bash
$ curl http://localhost:4000/v1/hostinfo | jq
{
  "id": "QmRmkUZCHZ1LdvayoKFtevmzE7RzkVVjGKA6uZ9yHoPCUW",
  "addresses": [
    "/ip4/127.0.0.1/tcp/46079/p2p/QmRmkUZCHZ1LdvayoKFtevmzE7RzkVVjGKA6uZ9yHoPCUW"
  ],
  "protocols": [
    "/p2p/id/delta/1.0.0",
    "/ipfs/id/1.0.0",
    "/ipfs/id/push/1.0.0",
    "/ipfs/ping/1.0.0",
    "/libp2p/grpc/1.0.0"
  ],
  "peers": []
}

```

## Contributing

PRs accepted.

## License

MIT ?? [Daniel Ribeiro](https://github.com/drgomesp)

