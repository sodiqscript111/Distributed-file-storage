package main

import (
	"distributed_file_storage/p2p"
	"fmt"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	// call method on the instance (tr), not the package
	if err := tr.ListenAndAccept(); err != nil {
		fmt.Println("There is an error", err)
	}

	select {} // keep the program alive
}
