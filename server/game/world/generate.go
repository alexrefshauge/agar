package world

import (
	"math"
	"math/rand"

	o "github.com/alexrefshauge/agar/server/game/object"
)

func (w *World) RegisterObject(o o.GameObject) bool {
	id := o.GetId()
	if _, notAvailable := w.IdMap[id]; notAvailable {
		//TODO: err with type of game object occupying the slot
		return false
	}

	w.IdMap[id] = o
	return true
}

func (w *World) NewBaseObject(x, y float32) o.BaseGameObject {
	id := w.NextId()
	return o.BaseGameObject{
		Id: id,
		Position: o.Vec{
			X: x,
			Y: y,
		}}
}

func (w *World) Generate(blobCount int) {
	for range blobCount {
		dir := rand.Float64()*2*math.Pi - math.Pi
		rad := rand.Float64()
		x, y :=
			float32(math.Cos(dir)*float64(w.Size)*rad),
			float32(math.Sin(dir)*float64(w.Size)*rad)
		w.AddBlob(x, y)
	}
}

func (w *World) NextId() int {
	for i := 0; ; i++ {
		if _, hasValue := w.IdMap[i]; !hasValue {
			return i
		}
	}
}
