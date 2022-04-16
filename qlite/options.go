package qlite

import (
	"bytes"

	"github.com/culionbear/go-qlite/network"
	"github.com/culionbear/qtool/ds/queue"
)

func (m *Manager) Do(cmd... []byte) (*queue.Manager[any], error) {
	writer := &bytes.Buffer{}
	for _, v := range cmd {
		writer.Write(v)
	}
	ch := make(chan *queue.Manager[any])
	err := m.client.Write(&network.Message{
		Buf: writer.Bytes(),
		Ch: ch,
	})
	if err != nil {
		return nil, err
	}
	return <- ch, nil
}

func (m *Manager) FPack(cmd... any) Func {
	writer := &bytes.Buffer{}
	for _, v := range cmd {
		writer.Write(m.handler.Pack(v))
	}
	return m.handler.PackFunc(writer.Bytes())
}

func (m *Manager) Pack(cmd... any) []byte {
	length := len(cmd)
	list := make([][]byte, length)
	for i := 0; i < length; i ++ {
		switch v := cmd[i].(type) {
		case Func:
			list[i] = v
		default:
			list[i] = m.handler.Pack(cmd[i])
		}
	}
	return nil
}
