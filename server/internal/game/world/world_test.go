package world

import (
	"testing"

	"github.com/alexrefshauge/agar/server/internal/game/object"
	"github.com/alexrefshauge/agar/server/pkg/vec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorld(t *testing.T) {
	world := New(100)
	blob := object.NewBlob(world.NewId(), vec.NewVec2(0, 0), 1)

	otherSameId := object.NewBlob(blob.Id(), vec.NewVec2(42, 69), 4)

	ok := world.AddObject(blob)
	assert.Equal(t, 1, len(world.objects))
	require.True(t, ok, "should add object")

	assert.Equal(t, 1, len(world.GetObjects()))

	ok = world.AddObject(otherSameId)
	assert.Equal(t, 1, len(world.objects))
	require.False(t, ok, "should not add object")

	assert.Equal(t, 1, len(world.GetObjects()))

	ok = world.RemoveObject(blob.Id() + 1)
	assert.Equal(t, 1, len(world.objects))
	require.False(t, ok, "should not delete non-existing object")

	assert.Equal(t, 1, len(world.GetObjects()))

	ok = world.RemoveObject(blob.Id())
	assert.Equal(t, 0, len(world.objects))
	require.True(t, ok, "should delete existing object")
}
