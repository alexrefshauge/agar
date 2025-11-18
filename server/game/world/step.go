package world

import (
	"time"
)

// Step
//
// dt: delta time in seconds
func (w *World) Step() {
	deltaTime := time.Since(w.lastUpdate)

	objects := w.IdMap
	for _, o := range objects {
		if o.Update(deltaTime.Seconds()) {
			w.updates <- o.GetId()
		}
		deltaTime = time.Since(w.lastUpdate)

	}

	w.lastUpdate = time.Now()
}
