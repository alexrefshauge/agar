package network

import (
	"github.com/alexrefshauge/agar/server/internal/game/iface"
	"github.com/alexrefshauge/agar/server/internal/game/object"
)

type Packet interface {
	Serialize() []byte
}

type packetImpl struct{}

func (p *packetImpl) Serialize() []byte { panic("packet is not serializable") }

type WelcomePacket struct {
	packetImpl
	clientId int
	playerId int
}

type StatePacket struct {
	packetImpl
	Players []*object.Player
	Blobs   []*object.Blob
}

type DeltaStatePacket struct {
	packetImpl
	Load   []iface.Object
	Unload []int
	Players []*object.Player
	Blobs   []*object.Blob
}

type PlayerInputPacket struct {
	packetImpl
	ClientId  int
	Direction float64
}

const (
	WELCOME uint8 = iota + 1
	STATE
	DELTA_STATE
	PLAYER_INPUT
)
