package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"

	"github.com/alexrefshauge/agar/server/game/object"
)

type updateJson struct {
	Players []*object.Player `json:"players"`
	Blobs   []*object.Blob   `json:"blobs"`
}

func (g *Game) broadcastUpdates() {

	clients := g.clients
	updates := g.GetPendingUpdates()
	players := make([]*object.Player, 0)
	blobs := make([]*object.Blob, 0)
	for _, o := range updates {
		if blob, ok := o.(*object.Blob); ok {
			blobs = append(blobs, blob)
		}
		if player, ok := o.(*object.Player); ok {
			players = append(players, player)
		}
	}

	if len(updates) == 0 {
		slog.Debug("no world updates", "tick start", g.tickStart)
		return
	}
	// packet with all data
	data, err := json.Marshal(updateJson{Players: players, Blobs: blobs})
	packet := make([]byte, PACKET_METADATA_SIZE+len(data))
	copy(packet[0:len(data)], data)
	copy(packet[len(data):], PACKET_STOP)

	for id, client := range clients {
		fmt.Printf("sending updates to client %d\n%s\n", id, string(packet))
		_, err = client.Write(packet)
		if errors.Is(err, net.ErrClosed) {
			fmt.Printf("Client %d disconnected\n", id)
			g.removeClient(id)
			continue
		}
		if err != nil {
			fmt.Printf("Failed to send data to client %d\n", id)
		}
		fmt.Printf("%d bytes sent to client %d\n", len(packet), id)
	}
}
