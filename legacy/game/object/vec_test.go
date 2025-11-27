package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosition(t *testing.T) {
	p1 := Vec{
		X: 0, Y: 0}
	p1.Move(42, 69)

	expected := Vec{42, 69}

	assert.Equal(t, expected.X, p1.X)
	assert.Equal(t, expected.Y, p1.Y)
}

func TestDistance(t *testing.T) {
	var actual, expected float32 = 0, 0

	p1 := Vec{
		X: 42, Y: 69}

	p2 := Vec{
		X: 43, Y: 69}

	p3 := Vec{
		X: 41, Y: 68}

	actual = p1.DistanceToPoint(&p2)
	expected = 1
	assert.Equal(t, expected, actual)

	actual = p1.DistanceToPoint(&p3)
	expected = 1
	assert.Greater(t, actual, expected)
}

func TestNorm(t *testing.T) {

}
