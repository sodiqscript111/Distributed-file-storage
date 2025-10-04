package p2p

// handshakeFunc is
type HandshakeFunc func(Peer) error

func NOPhandshake(Peer) error {
	return nil
}
