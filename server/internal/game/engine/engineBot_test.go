package engine

import (
	"math"
	"testing"

	"github.com/alexrefshauge/agar/server/internal/game/object"
	"github.com/alexrefshauge/agar/server/pkg/vec"
	"github.com/stretchr/testify/assert"
)

var tolerance = 1e-6

func Test_BotDirFunc(t *testing.T) {
	engine := New()
	bot := object.NewPlayer(engine.world.NewId(), vec.NewVec2(0, 0), "bot", 1, botDirFunc)
	engine.world.AddObject(bot)
	blob := object.NewBlob(engine.world.NewId(), vec.NewVec2(1, 1), 1)
	engine.world.AddObject(blob)

	engine.Step(1)
	expected := math.Pi / 4
	actual := bot.Dir

	assert.InDelta(t, expected, actual, tolerance, "bot should point towards blob")
}
