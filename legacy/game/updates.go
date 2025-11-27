package game

import (
	"github.com/alexrefshauge/agar/server/game/object"
)

func (g *Game) GetPendingUpdates() []object.GameObject {
	if len(g.world.Updates) == 0 {
		return []object.GameObject{}
	}
	ids := make([]int, 0, 1024)
	updateId := <-g.world.Updates
	for updateId != -1 {
		ids = append(ids, updateId)
		updateId = <-g.world.Updates
	}

	objects := make([]object.GameObject, len(ids))
	for i, id := range ids {
		objects[i] = g.world.IdMap[id]
	}
	return objects
}
