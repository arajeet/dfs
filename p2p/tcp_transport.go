package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

// If the return type is TCP Transport then calling NewTCP Transport allows us
// to access the elements, but if it type of Transport Interface then itssef

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		//	defer conn.Close()
		//fmt.Println(conn.)
		if err != nil {
			fmt.Printf("TCP Accept err: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

type Test struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("new Incoming Connections %+v\n", peer)
	msg := &Test{}
	if err := t.HandshakeFunc(peer); err != nil {

		conn.Close()
		return

	}

	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error : %s\n", err)
			continue
		}
	}

}
