package game

import (
	"net"

	"github.com/alexrefshauge/agar/server/game/world"
)

type Game struct {
	world   world.World
	clients map[int]net.Conn // client map with player object id as key
}

func NewGameWithWorld(w world.World) *Game {
	return &Game{
		world:   w,
		clients: make(map[int]net.Conn),
	}
}

func NewGame() *Game {
	return &Game{
		world:   *world.NewWorld(),
		clients: make(map[int]net.Conn),
	}
}
