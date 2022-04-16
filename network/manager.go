package network

import (
	"errors"
	"io"
	"net"

	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/protocol"
)

type Manager struct {
	password	[]byte
	conn		net.Conn
	ch			chan chan *queue.Manager[any]
	handler		*protocol.Manager
}

func New(addr, pass string, cap int) (*Manager, error) {
	password := []byte(pass)
	for i := len(pass); i < 256; i ++ {
		password = append(password, 0x00)
	}
	conn, err := net.Dial("tcp", addr)
	return &Manager{
		password: password,
		conn: conn,
		handler: protocol.New(),
		ch: make(chan chan *queue.Manager[any], cap),
	}, err
}

func (m *Manager) Run() error {
	defer m.Close()
	m.conn.Write(m.password)
	if err := m.verify(); err != nil {
		return err
	}
	return m.read()
}

func (m *Manager) Close() error {
	return m.conn.Close()
}

func (m *Manager) verify() error {
	buf := make([]byte, 3, 3)
	_, err := io.ReadAtLeast(m.conn, buf, 3)
	if err != nil {
		return err
	}
	if buf[2] != 0x04 {
		return errors.New("password is not true")
	}
	return nil
}

func (m *Manager) read() error {
	request := make([]byte, 1024)
	for {
		n, err := m.conn.Read(request)
		if err != nil {
			return err
		}
		msg := request[:n]
		for {
			size, point, isFinish := m.handler.PackSize(msg)
			if !isFinish {
				break
			}
			info, err := m.handler.Read(m.conn, size, msg[point:])
			if err != nil {
				return err
			}
			cmd, qerr := m.handler.Unpack(info)
			if qerr != nil {
				return qerr
			}
			<- m.ch <- cmd
			if size + point >= n {
				break
			}
			msg = msg[size + point:]
		}
	}
}

func (m *Manager) Write(msg *Message) error {
	_, err := m.conn.Write(m.handler.Write(msg.Buf))
	if err != nil {
		return err
	}
	m.ch <- msg.Ch
	return nil
}
