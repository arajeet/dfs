package p2p

import "errors"

// ErrInvalidHanshake is returned if the handshake between
// local and remote node could not be established
var ErrInvalidHandshake = errors.New("Invalid Handshake")

// Handshake func is ?
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
