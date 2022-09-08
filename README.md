# go-libp2p-hostinfo

[![madeby](https://img.shields.io/badge/made%20by-%40drgomesp-blue)](https://github.com/drgomesp/)
[![Go Report Card](https://goreportcard.com/badge/github.com/drgomesp/go-libp2p-grpc)](https://goreportcard.com/report/github.com/drgomesp/go-libp2p-grpc)
[![build](https://github.com/drgomesp/go-libp2p-grpc/actions/workflows/go-test.yml/badge.svg?style=squared)](https://github.com/drgomesp/go-libp2p-grpc/actions)
[![codecov](https://codecov.io/gh/drgomesp/go-libp2p-grpc/branch/main/graph/badge.svg?token=BRMFJRJV2X)](https://codecov.io/gh/drgomesp/go-libp2p-grpc)


> A pluggable libp2p host service that exposes general information about the host and the network.

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

## Table of Contents

- [Install](#install)
- [Features](#features)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Install

```bash
go get github.com/drgomesp/go-libp2p-grpc
```

## Features


## Usage

## Contributing

PRs accepted.

## License

MIT © [Daniel Ribeiro](https://github.com/drgomesp)

