package object

import (
	"github.com/alexrefshauge/agar/server/internal/game/iface"
	"github.com/alexrefshauge/agar/server/pkg/vec"
)

type Blob struct {
	object
	Size int `json:"size"`
}

func NewBlob(id int, p *vec.Vec2, size int) *Blob {
	return &Blob{
		object: New(id, p),
		Size:   size,
	}
}

func (b *Blob) Update(dt float64, world iface.World) bool {
	return false
}
