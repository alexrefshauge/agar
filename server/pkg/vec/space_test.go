package vec

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVec2_Distance(t *testing.T) {
	a := NewVec2(0, 0)
	b := NewVec2(1, 0)

	expected := 1
	actual := a.DistanceTo(b)

	assert.InDelta(t, expected, actual, tolerance)
}

func TestVec2_DistanceGreater(t *testing.T) {
	a := NewVec2(-1, -1)
	b := NewVec2(1, 1)

	actual := a.DistanceTo(b)
	than := 2.0

	assert.Greater(t, actual, than, tolerance)
}

func TestVec2_DistanceZero(t *testing.T) {
	a := NewVec2(42, 69)
	b := NewVec2(42, 69)

	actual := a.DistanceTo(b)
	expected := 0

	assert.InDelta(t, expected, actual, tolerance)
}

func TestVec2_Angle(t *testing.T) {

	actual90 := NewVec2(0, 1).Angle()
	expected90 := math.Pi / 2

	actual135 := NewVec2(-1, 1).Angle()
	expected135 := math.Pi * (3.0 / 4.0)

	actual270 := NewVec2(0, -1).Angle()
	expected270 := -(math.Pi / 2.0)

	assert.InDelta(t, expected90, actual90, tolerance)
	assert.InDelta(t, expected135, actual135, tolerance)
	assert.InDelta(t, expected270, actual270, tolerance)
}

func TestVec2_AngleTo(t *testing.T) {
	a := NewVec2(1, 0)
	b90 := NewVec2(0, 1)
	b45 := NewVec2(1, 1)
	bNeg90 := NewVec2(0, -1)

	actual90 := a.AngleTo(b90)
	actual45 := a.AngleTo(b45)
	actualNeg90 := a.AngleTo(bNeg90)

	require.InDelta(t, math.Pi/2, actual90, tolerance)
	require.InDelta(t, math.Pi/4, actual45, tolerance)
	require.InDelta(t, -math.Pi/2, actualNeg90, tolerance)
}
