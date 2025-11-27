package engine

import (
	"math/rand/v2"

	"github.com/alexrefshauge/agar/server/internal/game/object"
	"github.com/alexrefshauge/agar/server/pkg/vec"
)

func (e *Engine) randomPosition() *vec.Vec2 {
	angle := rand.Float64()
	dist := rand.Float64() * float64(e.world.Size())
	return vec.Vec2FromAngle(angle).MulScalar(dist)
}

func (e *Engine) Generate(count int) {
	for range count {
		size := 10 + rand.IntN(10)
		blob := object.NewBlob(e.world.NewId(), e.randomPosition(), size)
		e.world.AddObject(blob)

		e.world.AddObject(object.NewPlayer(e.world.NewId(), vec.NewVec2(0, 0), "bot", 30, botDirFunc))
	}
}
