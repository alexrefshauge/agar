package world

import (
	"time"

	"github.com/alexrefshauge/agar/server/game/object"
)

// Step
//
// dt: delta time in seconds
func (w *World) Step() {
	deltaTime := time.Since(w.lastUpdate)

	objects := w.IdMap
	for _, o := range objects {
		if o.Update(deltaTime.Seconds()) {
			w.Updates <- o.GetId()
		}
		if player, ok := o.(*object.Player); ok {
			w.playerCollide(player)
		}
		deltaTime = time.Since(w.lastUpdate)
	}

	w.Updates <- -1
	w.lastUpdate = time.Now()
}

func (w *World) playerCollide(player *object.Player) {
	distanceOutside := (player.Position.Len() + float32(player.Size)) - w.Size
	if distanceOutside < 0 {
		return
	}

	diff := player.Position.Norm().Scale(-1 * distanceOutside)

	player.Position = *player.Position.Add(diff)
}
