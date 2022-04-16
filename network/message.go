package network

import "github.com/culionbear/qtool/ds/queue"

type Message struct {
	Buf	[]byte
	Ch	chan *queue.Manager[any]
}
