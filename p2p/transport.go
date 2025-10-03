package p2p

// Peer is an interface that represent the remote note
type Peer interface {
}

//transport is anything that handles communication
//between the nodes in the network, This can be (TCP, UDP , websocket ...)

type Transport interface {
	listenAndAccept() error
}
