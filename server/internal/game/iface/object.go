package iface

import "github.com/alexrefshauge/agar/server/pkg/vec"

type Object interface {
	Update(float64, World) bool
	Id() int
	Pos() *vec.Vec2
}
