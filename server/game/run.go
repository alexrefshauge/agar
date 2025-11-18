package game

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func (g *Game) Run() {
	go g.runBroadcastLoop()
	for {
		g.world.Step()
		time.Sleep(500)
	}
}

func (g *Game) Listen(address string) {
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		panic("failed to open server socket")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("failed to accept client socket")
		}
		id := g.AddClient(conn)

		packet := fmt.Appendf([]byte{}, "%d\r\n", id)

		w := g.world
		worldData, err := json.Marshal(w)
		packet = append(packet, worldData...)
		packet = append(packet, PACKET_STOP...)
		if err != nil {
			g.RemoveClient(id)
			continue
		}
		_, err = conn.Write(packet)
		if err != nil {
			g.RemoveClient(id)
			fmt.Printf("Failed to add client %d\n", id)
			continue
		}
	}
}

func (g *Game) AddClient(conn net.Conn) int {
	p := g.world.AddPlayer()
	id := p.GetId()
	g.clients[id] = conn
	go g.handleClient(id, conn)
	return id
}

func (g *Game) RemoveClient(id int) {
	client, ok := g.clients[id]
	if !ok {
		fmt.Printf("Failed to remove client. Client %d does not exist\n", id)
		return
	}

	delete(g.clients, id)
	client.Close()
}
