package vec

import (
	"math"
)

func (v *Vec2) DistanceTo(other *Vec2) float64 {
	return other.Sub(v).Norm()
}

func (v *Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v *Vec2) AngleTo(other *Vec2) float64 {
	dot := v.X*other.X + v.Y*other.Y
	det := v.X*other.Y - v.Y*other.X
	vNorm, otherNorm := v.Norm(), other.Norm()
	if vNorm == 0 || otherNorm == 0 {
		return 0
	}

	return math.Atan2(det, dot)
}
