package object

import (
	"math"
)

var speed float32 = 100

func (p *Player) Update(deltaTime float64) bool {
	dt := float32(deltaTime)
	p.Velocity.X = speed * float32(math.Cos(float64(p.Direction)))
	p.Velocity.Y = speed * float32(math.Sin(float64(p.Direction)))
	p.Position.Move(p.Velocity.X*dt, p.Velocity.Y*dt)

	return true
}

func (b *Blob) Update(dt float64) bool {
	return false
}
