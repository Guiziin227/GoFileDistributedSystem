package main

import (
	"log"

	"github.com/Guiziin227/GoFileDistributedSystem/p2p"
)

func main() {

	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to start TCP transport: %v", err)
	}

	select {}
}
