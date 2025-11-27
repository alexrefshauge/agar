package network

import (
	"math/rand/v2"
	"testing"

	"github.com/alexrefshauge/agar/server/internal/game/object"
	"github.com/alexrefshauge/agar/server/pkg/vec"
	"github.com/stretchr/testify/assert"
)

func TestPacket(t *testing.T) {
	a, b := rand.IntN(254), rand.IntN(254)
	actual := (&WelcomePacket{
		clientId: a,
		playerId: b,
	}).Serialize()

	expected := []byte{
		0b_0000_0001,
		0b_0000_0000, 0b_0000_0000, 0b_0000_0000, byte(a),
		0b_0000_0000, 0b_0000_0000, 0b_0000_0000, byte(b),
	}

	assert.Equal(t, expected, actual)
}

func TestStatePacket(t *testing.T) {
	actual := (&StatePacket{
		Players: []*object.Player{
			object.NewPlayer(2, vec.NewVec2(1, 2), "hello", 1, nil),
			object.NewPlayer(1, vec.NewVec2(3, 4), "someone", 200, nil),
			object.NewPlayer(3, vec.NewVec2(5, 6), "123456789abcdef", 300, nil),
		},
		Blobs: []*object.Blob{
			object.NewBlob(0, vec.NewVec2(1, 1), 42),
			object.NewBlob(4, vec.NewVec2(1, 1), 42),
			object.NewBlob(5, vec.NewVec2(1, 1), 42),
			object.NewBlob(6, vec.NewVec2(1, 1), 42),
			object.NewBlob(7, vec.NewVec2(1, 1), 42),
		},
	}).Serialize()

	expected := []byte{
		byte(STATE),
	}

	assert.Equal(t, expected, actual)
}
