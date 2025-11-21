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
	Unload  []int            `json:"unload"`
	Eat     []int            `json:"eat"`
}

func wrap(d updateJson) []byte {
	data, err := json.Marshal(d)
	if err != nil {
		slog.Error("failed to marshal updateJson", "err", err)
		return []byte{}
	}
	packet := make([]byte, PACKET_METADATA_SIZE+len(data))
	copy(packet[0:len(data)], data)
	copy(packet[len(data):], PACKET_STOP)
	return packet
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

	for id, client := range clients {
		iPlayers, iBlobs := g.getPlayersAndBlobsByInterest(client.interests)
		updatePacket := updateJson{
			Players: iPlayers,
			Blobs:   iBlobs,
			Unload:  g.updatePlayerClientInterests(id),
			Eat:     g.world.JustEaten,
		}

		if len(updatePacket.Eat) > 0 {
			slog.Debug("blobs eaten", "count", len(updatePacket.Eat))
		}

		// packet with all data
		packet := wrap(updatePacket)

		//fmt.Printf("sending updates to client %d\n%s\n", id, string(packet))
		_, err := client.conn.Write(packet)
		if errors.Is(err, net.ErrClosed) {
			fmt.Printf("Client %d disconnected\n", id)
			g.removeClient(id)
			continue
		}
		if err != nil {
			fmt.Printf("Failed to send data to client %d\n", id)
		}
	}
}

var loadRadius float32 = 1000

func (g *Game) updatePlayerClientInterests(id int) []int {
	var player *object.Player
	client, online := g.clients[id]
	if !online {
		return nil
	}
	pObject, inWorld := g.world.IdMap[id]
	if !inWorld {
		return nil
	}

	player, isPlayer := pObject.(*object.Player)
	if !isPlayer {
		return nil
	}

	unloadIds := make([]int, 0, 8)
	for _, o := range g.world.IdMap {
		objectId := o.GetId()
		if objectId == player.Id {
			client.interests[objectId] = true
		}

		objectPos := o.GetPos()
		dist := player.Position.Sub(&objectPos)
		load := dist.Len() < loadRadius

		isLoaded := client.interests[objectId]
		if load {
			if isLoaded {
				continue
			}

			client.interests[objectId] = true
			continue
		}

		if isLoaded {
			client.interests[objectId] = false
			unloadIds = append(unloadIds, objectId)
			continue
		}
	}

	return unloadIds
}

func filterObjectsByInterest[T object.GameObject](interests map[int]bool, objects []T) []T {
	filtered := make([]T, 0, len(objects))
	for _, object := range objects {
		if interests[object.GetId()] {
			filtered = append(filtered, object)
		}
	}
	return filtered
}

func (g *Game) getPlayersAndBlobsByInterest(interests map[int]bool) ([]*object.Player, []*object.Blob) {
	p := filterObjectsByInterest[*object.Player](interests, g.world.Players)
	o := filterObjectsByInterest[*object.Blob](interests, g.world.Blobs)
	return p, o
}
