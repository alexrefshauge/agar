package game

import (
	"fmt"
	"net"
	"strconv"

	"github.com/alexrefshauge/agar/server/game/object"
)

func (g *Game) handleClient(id int, conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4096)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Failed to receive packet: %s\n", err)
			return
		}
		packet := buffer[:n]
		fmt.Printf("received %s from client %d\n", string(packet), id)

		direction, err := strconv.ParseFloat(string(packet), 64)
		fmt.Println(direction)
		o, ok := g.world.IdMap[id]
		if !ok {
			//TODO: handle gracefully
			fmt.Printf("Player with client id %d does not exist. Closing connection", id)
			conn.Close()
			return
		}

		switch object := o.(type) {
		case *object.Player:
			object.Direction = float32(direction)
			break
		default:
			fmt.Println("GameObject id and client id mismatch. GameObject is not a Player")
		}
	}
}
