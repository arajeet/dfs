package main

import (
	"github/arajeet/distributedfs/p2p"
	"log"
)

func main() {
	opts := p2p.TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(opts)

	err := tr.ListenAndAccept()
	if err != nil {
		log.Fatal(err)
	}
	select {}
	//
	// fmt.Println("Starting ")
}
