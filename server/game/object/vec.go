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
