package qlite

import (
	"github.com/culionbear/go-qlite/network"
	"github.com/culionbear/qtool/protocol"
)

type Func []byte

type Manager struct {
	client	*network.Manager
	handler	*protocol.Manager
}

func New(conf *Config) (*Manager, error) {
	c, err := network.New(
		conf.Addr,
		conf.Password,
		conf.Cap,
	)
	if err != nil {
		return nil, err
	}
	go c.Run()
	return &Manager{
		client: c,
		handler: protocol.New(),
	}, nil
}
