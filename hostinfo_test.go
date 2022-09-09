package hostinfo_test

import (
	"context"
	"testing"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"

	hostinfo "github.com/drgomesp/go-libp2p-hostinfo"
)

func newHost(t *testing.T, listen multiaddr.Multiaddr) host.Host {
	h, err := libp2p.New(
		libp2p.ListenAddrs(listen),
	)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

func TestNewService(t *testing.T) {
	ma, _ := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/10000")
	h := newHost(t, ma)
	defer h.Close()

	svc, err := hostinfo.NewService(context.Background(), h)

	assert.NoError(t, err)
	assert.NotNil(t, svc)
}
