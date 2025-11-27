package vec

import "math"

type Vec2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewVec2(x, y float64) *Vec2 {
	return &Vec2{X: x, Y: y}
}
func Vec2FromAngle(a float64) *Vec2 {
	return &Vec2{X: math.Cos(a), Y: math.Sin(a)}
}

type op func(x float64) float64

func (v *Vec2) applyOp(f op) *Vec2 {
	return &Vec2{X: f(v.X), Y: f(v.Y)}
}

func (v *Vec2) AddScalar(s float64) *Vec2 { return v.applyOp(func(x float64) float64 { return x + s }) }
func (v *Vec2) SubScalar(s float64) *Vec2 { return v.applyOp(func(x float64) float64 { return x - s }) }
func (v *Vec2) MulScalar(s float64) *Vec2 { return v.applyOp(func(x float64) float64 { return x * s }) }
func (v *Vec2) DivScalar(s float64) *Vec2 { return v.applyOp(func(x float64) float64 { return x / s }) }

type vecOp func(a, b float64) float64

func (v *Vec2) applyVecOp(f vecOp, other *Vec2) *Vec2 {
	return &Vec2{X: f(v.X, other.X), Y: f(v.Y, other.Y)}
}

func (v *Vec2) Add(other *Vec2) *Vec2 {
	return v.applyVecOp(func(a, b float64) float64 { return a + b }, other)
}
func (v *Vec2) Sub(other *Vec2) *Vec2 {
	return v.applyVecOp(func(a, b float64) float64 { return a - b }, other)
}
func (v *Vec2) Mul(other *Vec2) *Vec2 {
	return v.applyVecOp(func(a, b float64) float64 { return a * b }, other)
}
func (v *Vec2) Div(other *Vec2) *Vec2 {
	return v.applyVecOp(func(a, b float64) float64 { return a / b }, other)
}

// Norm returns the length of the vector
func (v *Vec2) Norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalized returns the unit vector of this vector
func (v *Vec2) Normalized() *Vec2 {
	norm := v.Norm()
	return &Vec2{X: v.X / norm, Y: v.Y / norm}
}
