package world

import (
	"time"

	o "github.com/alexrefshauge/agar/server/game/object"
)

type World struct {
	Players    []*o.Player `json:"players"`
	Blobs      []*o.Blob   `json:"blobs"`
	IdMap      map[int]o.GameObject
	lastUpdate time.Time

	Size float32 `json:"size"` // radius of the world border

	Updates chan int
}

func NewWorld() *World {
	return &World{
		make([]*o.Player, 0),
		make([]*o.Blob, 0),
		make(map[int]o.GameObject),

		time.Now(),
		1000,

		make(chan int, 4096),
	}
}
