package game

import (
	"github.com/alexrefshauge/agar/server/game/object"
)

func (g *Game) GetPendingUpdates() []object.GameObject {
	ids := g.world.Flush()
	objects := make([]object.GameObject, len(ids))
	for i, id := range ids {
		objects[i] = g.world.IdMap[id]
	}
	return objects
}
