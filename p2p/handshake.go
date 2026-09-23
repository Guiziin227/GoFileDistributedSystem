package p2p

// HandshakeFunc é uma função que realiza o handshake entre dois nós.
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}
