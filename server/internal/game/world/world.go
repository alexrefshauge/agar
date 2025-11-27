package world

import (
	"github.com/alexrefshauge/agar/server/internal/game/iface"
	"github.com/alexrefshauge/agar/server/internal/game/object"
)

type World struct {
	objects map[int]iface.Object
	size    int
	players []object.Player
	removed []int
}

var _ iface.World = (*World)(nil)

func (w *World) Size() int                     { return w.size }
func (w *World) Objects() map[int]iface.Object { return w.objects }

func New(size int) *World {
	return &World{
		objects: make(map[int]iface.Object, 0),
		size:    size,
		removed: make([]int, 0),
	}
}

func (w *World) NewId() int {
	for id := 0; ; id++ {
		if _, ok := w.objects[id]; !ok {
			return id
		}
	}
}

func (w *World) Cleanup() {
	w.removed = w.removed[:0]
}

func (w *World) AddObject(o iface.Object) bool {
	if _, ok := w.objects[o.Id()]; ok {
		return false
	}
	w.objects[o.Id()] = o
	return true
}

func (w *World) RemoveObject(id int) bool {
	if _, ok := w.objects[id]; ok {
		delete(w.objects, id)
		w.removed = append(w.removed, id)
		return true
	}
	return false
}

func (w *World) GetRemovals() []int {
	return w.removed
}

func (w *World) GetObjects() map[int]iface.Object {
	return w.objects
}
