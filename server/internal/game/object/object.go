package object

import (
	"github.com/alexrefshauge/agar/server/pkg/vec"
)

type object struct {
	id  int       `json:"id"`
	pos *vec.Vec2 `json:"pos"`
}

func New(id int, p *vec.Vec2) object {
	return object{
		id:  id,
		pos: p,
	}
}

func (o *object) Id() int {
	return o.id
}

func (o *object) Pos() *vec.Vec2 {
	return &vec.Vec2{X: o.pos.X, Y: o.pos.Y}
}
