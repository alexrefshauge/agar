package world

import (
	"log/slog"
	"time"

	"github.com/alexrefshauge/agar/server/game/object"
)

// Step
//
// dt: delta time in seconds
func (w *World) Step() {
	w.JustEaten = w.JustEaten[:0]
	w.handleEatPlayers()

	deltaTime := time.Since(w.lastUpdate)

	objects := w.IdMap
	for _, o := range objects {
		if o.Update(deltaTime.Seconds()) {
			w.Updates <- o.GetId()
		}
		if blob, ok := o.(*object.Blob); ok {
			w.handleEat(blob)
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

func (w *World) handleEat(blob *object.Blob) {
	for _, player := range w.Players {
		dist := player.Position.DistanceToPoint(&blob.Position)
		eatDist := float32(player.Size) + float32(blob.Size)
		if dist < eatDist {
			player.Size += (blob.Size / 10) //blob.Size
			w.Remove(blob)
			w.JustEaten = append(w.JustEaten, blob.GetId())
			slog.Debug("blog eaten", "blob id", blob.GetId())
		}
	}
}

func (w *World) handleEatPlayers() {
	toRemove := make([]*object.Player, 0)
	for i, playerA := range w.Players {
		for j, playerB := range w.Players {
			// Check if player A shoud be eaten by player B
			if i == j {
				continue
			}

			dist := playerA.Position.DistanceToPoint(&playerB.Position)
			if dist < float32(playerB.Size) {
				playerB.Size += playerA.Size / 10
				toRemove = append(toRemove, playerA)
			}
		}
	}

	for _, p := range toRemove {
		w.JustEaten = append(w.JustEaten, p.GetId())
		w.Remove(p)
	}
}
