package game

import (
	"net"
	"time"

	"github.com/alexrefshauge/agar/server/game/world"
)

type Game struct {
	world     world.World
	tickStart time.Time

	clients     map[int]net.Conn // client map with player object id as key
	clientQueue chan net.Conn
}

func NewGameWithWorld(w world.World) *Game {
	return &Game{
		world:       w,
		clients:     make(map[int]net.Conn),
		clientQueue: make(chan net.Conn, 64),
	}
}

func NewGame() *Game {
	return &Game{
		world:       *world.NewWorld(),
		clients:     make(map[int]net.Conn),
		clientQueue: make(chan net.Conn, 64),
	}
}
