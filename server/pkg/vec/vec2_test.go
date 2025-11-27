package vec

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tolerance = 1e-6

func assertEqualVec2(t *testing.T, expected, actual *Vec2) {
	assert.InDelta(t, expected.X, actual.X, tolerance)
	assert.InDelta(t, expected.Y, actual.Y, tolerance)
}

func TestVec2_AddScalar(t *testing.T) {
	actual := (&Vec2{1, 2}).AddScalar(1.5)
	expected := Vec2{2.5, 3.5}

	assertEqualVec2(t, &expected, actual)
}

func TestVec2_SubScalar(t *testing.T) {
	actual := (&Vec2{2, 3}).SubScalar(0.5)
	expected := Vec2{1.5, 2.5}

	assertEqualVec2(t, &expected, actual)
}

func TestVec2_Add(t *testing.T) {
	actual := (&Vec2{2, 3}).Add(&Vec2{1, 2})
	expected := Vec2{3, 5}

	assertEqualVec2(t, &expected, actual)
}

func TestVec2_Mul(t *testing.T) {
	actual := (&Vec2{2, 3}).Mul(&Vec2{0.5, 2.5})
	expected := Vec2{1, 7.5}

	assertEqualVec2(t, &expected, actual)
}
func TestVec2_Div(t *testing.T) {
	actual := (&Vec2{2, 4}).Div(&Vec2{2, 0.5})
	expected := Vec2{1, 8}

	assertEqualVec2(t, &expected, actual)
}

// Other operations
func TestVec2_Norm_1(t *testing.T) {
	actual := (&Vec2{1, 0}).Norm()
	expected := 1

	assert.InDelta(t, expected, actual, tolerance)
}

func TestVec2_Norm_1And1(t *testing.T) {
	actual := (&Vec2{1, 1}).Norm()
	expected := math.Sqrt(2)

	assert.InDelta(t, expected, actual, tolerance)
}

func TestVec2_Norm_1And1Negative(t *testing.T) {
	actual := (&Vec2{-1, -1}).Norm()
	expected := math.Sqrt(2)

	assert.InDelta(t, expected, actual, tolerance)
}
func TestVec2_Norm_1And1PartNegative(t *testing.T) {
	actual := (&Vec2{-1, 1}).Norm()
	expected := math.Sqrt(2)

	assert.InDelta(t, expected, actual, tolerance)
}

func TestVec2_NormalizedNormOne(t *testing.T) {
	actual := (&Vec2{1, 1}).Normalized().Norm()
	expected := 1

	assert.InDelta(t, expected, actual, tolerance)
}
func TestVec2_NormalizedOne(t *testing.T) {
	actual := (&Vec2{1, 0}).Normalized()
	expected := &Vec2{1, 0}

	assertEqualVec2(t, expected, actual)
}
func TestVec2_NormalizedTwo(t *testing.T) {
	actual := (&Vec2{2, 0}).Normalized()
	expected := &Vec2{1, 0}

	assertEqualVec2(t, expected, actual)
}

func TestVec2_NormalizedTwoNegative(t *testing.T) {
	actual := (&Vec2{-2, 0}).Normalized()
	expected := &Vec2{-1, 0}

	assertEqualVec2(t, expected, actual)
}
