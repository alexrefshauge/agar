package world

import (
	"log/slog"

	"github.com/alexrefshauge/agar/server/game/object"
)

func (w *World) Remove(o object.GameObject) {
	delete(w.IdMap, o.GetId())
	switch o.(type) {
	case *object.Player:
		playerIndex := -1
		for i, p := range w.Players {
			if p.GetId() == o.GetId() {
				playerIndex = i
				break
			}
		}

		//swap
		last := len(w.Players) - 1
		w.Players[playerIndex] = w.Players[last]
		w.Players[last] = nil
		w.Players = w.Players[:last]
		break

	case *object.Blob:
		blobIndex := -1
		for i, b := range w.Blobs {
			if b.GetId() == o.GetId() {
				blobIndex = i
				break
			}
		}

		//swap
		last := len(w.Blobs) - 1
		w.Blobs[blobIndex] = w.Blobs[last]
		w.Blobs[last] = nil
		w.Blobs = w.Blobs[:last]
		break

	default:
		slog.Warn("remove called on unknown game object type")
	}
}
