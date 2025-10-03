package p2p

// handshakeFunc is
type HandshakeFunc func(any) error

func NOPhandshake(Peer) error {
	return nil
}
