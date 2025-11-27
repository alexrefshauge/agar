package world

import (
	"math"
	"math/rand/v2"

	o "github.com/alexrefshauge/agar/server/game/object"
)

const PLAYER_START_SIZE = 100
const BLOB_START_SIZE = 10

func (w *World) AddBlob(x, y float32) *o.Blob {
	blob := &o.Blob{
		BaseGameObject: w.NewBaseObject(x, y),
		Size:           BLOB_START_SIZE,
	}

	w.RegisterObject(blob)
	w.Blobs = append(w.Blobs, blob)

	return blob
}

func (w *World) AddPlayer() *o.Player {
	base := w.NewBaseObject(0, 0)
	player := &o.Player{
		BaseGameObject: base,
		Name:           "new client",
		Size:           PLAYER_START_SIZE,
		Direction:      math.Pi - (rand.Float32() * math.Pi * 2),
	}
	w.RegisterObject(player)
	w.Players = append(w.Players, player)

	return player
}
