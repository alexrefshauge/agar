package engine

import (
	"math"

	"github.com/alexrefshauge/agar/server/internal/game/iface"
	"github.com/alexrefshauge/agar/server/internal/game/object"
)

var botDirFunc object.PlayerDirFunc = func(player *object.Player, world iface.World) float64 {
	var closest iface.Object
	var bestDist float64 = math.Inf(+1)
	for _, o := range world.GetObjects() {
		if o.Id() == player.Id() {
			continue
		}
		dist := player.Pos().DistanceTo(o.Pos())
		if dist < bestDist {
			bestDist = dist
			closest = o
		}
	}
	return closest.Pos().Sub(player.Pos()).Angle()
}

var dummyDirFunc object.PlayerDirFunc = func(player *object.Player, world iface.World) float64 {
	return player.Dir
}
