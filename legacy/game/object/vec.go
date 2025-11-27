package object

import "math"

type Vec struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func (p *Vec) Move(dx, dy float32) {
	p.X += dx
	p.Y += dy
}

func (p *Vec) DistanceToPoint(other *Vec) float32 {
	dx := float64(other.X - p.X)
	dy := float64(other.Y - p.Y)
	sum := math.Pow(dx, 2) + math.Pow(dy, 2)
	return float32(math.Sqrt(sum))
}

func (v *Vec) Len() float32 {
	return float32(math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2)))
}

func (v *Vec) Add(other *Vec) *Vec {
	return &Vec{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *Vec) Sub(other *Vec) *Vec {
	return &Vec{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v *Vec) Mul(other *Vec) *Vec {
	return &Vec{
		X: v.X * other.X,
		Y: v.Y * other.Y,
	}
}

func (v *Vec) Scale(s float32) *Vec {
	return &Vec{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// Div divides the numerator (self) by the denominator (other vector)
func (num *Vec) Div(den *Vec) *Vec {
	return &Vec{
		X: num.X / den.X,
		Y: num.Y / den.Y,
	}
}

func (v *Vec) Norm() *Vec {
	mag := v.Len()
	return &Vec{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}
