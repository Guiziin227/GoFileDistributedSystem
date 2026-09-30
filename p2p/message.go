package p2p

import "net"

// Message representa uma mensagem enviada entre nós na rede P2P.
type Message struct {
	From    net.Addr
	Payload []byte
}
