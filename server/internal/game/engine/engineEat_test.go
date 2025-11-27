package engine

import (
	"testing"

	"github.com/alexrefshauge/agar/server/internal/game/object"
	"github.com/alexrefshauge/agar/server/pkg/vec"
	"github.com/stretchr/testify/assert"
)

func Test_BotEat(t *testing.T) {
	engine := New()
	bot := object.NewPlayer(engine.world.NewId(), vec.NewVec2(0, 0), "bot", 2, botDirFunc)
	engine.world.AddObject(bot)
	blob := object.NewBlob(engine.world.NewId(), vec.NewVec2(1, 0), 1)
	engine.world.AddObject(blob)

	_, blobExists := engine.world.GetObjects()[blob.Id()]
	assert.True(t, blobExists, "blob should exist before step")

	engine.Step(1)

	_, blobExists = engine.world.GetObjects()[blob.Id()]
	assert.False(t, blobExists, "blob should be gone after step")
}

func Test_PlayerEat(t *testing.T) {
	engine := New()
	p1 := object.NewPlayer(engine.world.NewId(), vec.NewVec2(0, 0), "1", 3, botDirFunc)
	engine.world.AddObject(p1)
	p2 := object.NewPlayer(engine.world.NewId(), vec.NewVec2(1, 0), "2", 1, botDirFunc)
	engine.world.AddObject(p2)

	_, ok := engine.world.GetObjects()[p1.Id()]
	assert.True(t, ok, "player 1 should exist before step")
	_, ok = engine.world.GetObjects()[p2.Id()]
	assert.True(t, ok, "player 2 should exist before step")

	engine.Step(1)

	_, ok = engine.world.GetObjects()[p1.Id()]
	assert.True(t, ok, "player 1 should still exist after step")
	_, ok = engine.world.GetObjects()[p2.Id()]
	assert.False(t, ok, "player 2 should not exist after step")
}
