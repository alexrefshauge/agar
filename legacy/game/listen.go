package game

import (
	"fmt"
	"net"
)

func (g *Game) Listen(address string) {
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		panic("failed to open server socket")
	}

	for {
		conn, err := listener.Accept()
		fmt.Println("new connection")
		if err != nil {
			panic("failed to accept client socket")
		}
		g.clientQueue <- conn
	}
}
