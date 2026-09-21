package p2p

// Peer é uma interface que representa um nó remoto.
type Peer interface{}

// Transport é qualquer coisa que lide com a
// comunicação entre os nós no network.
// Isso pode ser na forma de (tcp, udp, websocket, etc).
type Transport interface {
	ListenAndAccept() error
}
