package game

import (
	"net"
	"time"

	"github.com/alexrefshauge/agar/server/game/world"
)

type Client struct {
	interests map[int]bool
	conn      net.Conn
}

type Game struct {
	world     world.World
	tickStart time.Time

	clients     map[int]Client
	clientQueue chan net.Conn
}

func NewGameWithWorld(w world.World) *Game {
	return &Game{
		world:       w,
		clients:     make(map[int]Client),
		clientQueue: make(chan net.Conn, 64),
	}
}

func NewGame() *Game {
	return &Game{
		world:       *world.NewWorld(),
		clients:     make(map[int]Client),
		clientQueue: make(chan net.Conn, 64),
	}
}
