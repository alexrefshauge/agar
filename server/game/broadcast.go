package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"
)

func (g *Game) runBroadcastLoop() {
	for {
		time.Sleep(200 * time.Millisecond)
		g.broadcastUpdates()
	}
}

func (g *Game) broadcastUpdates() {
	if len(g.clients) > 0 {
		fmt.Printf("broadcasting to %d clients\n", len(g.clients))
	}

	updates := g.GetPendingUpdates()

	clients := g.clients
	for id, client := range clients {
		data, err := json.Marshal(updates)
		fmt.Printf("\n\n\n%s\n\n\n", string(data))
		packet := make([]byte, PACKET_METADATA_SIZE+len(data))
		copy(packet[0:len(data)], data)
		copy(packet[len(data):], PACKET_STOP)
		_, err = client.Write(packet)
		if errors.Is(err, net.ErrClosed) {
			fmt.Printf("Client %d disconnected\n", id)
			g.RemoveClient(id)
			continue
		}
		if err != nil {
			fmt.Printf("Failed to send data to client %d\n", id)
		}
		fmt.Printf("%d bytes sent to client %d\n", len(packet), id)
	}
}
