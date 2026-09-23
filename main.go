package main

import (
	"log"

	"github.com/Guiziin227/GoFileDistributedSystem/p2p"
)

func main() {

	opts := p2p.TCPTransportOps{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       nil, // You can implement a decoder if needed
	}

	tr := p2p.NewTCPTransport(opts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to start TCP transport: %v", err)
	}

	select {}
}
