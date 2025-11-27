package object

import (
	"fmt"

	"github.com/alexrefshauge/agar/server/internal/game/iface"
	"github.com/alexrefshauge/agar/server/pkg/vec"
)

type PlayerDirFunc func(player *Player, world iface.World) float64

type Player struct {
	object
	Size int `json:"size"`

	Vel     *vec.Vec2 `json:"vel"`
	Dir     float64   `json:"dir"`
	Name    string    `json:"name"`
	dirFunc PlayerDirFunc
}

func NewPlayer(id int, p *vec.Vec2, name string, size int, dirFunc PlayerDirFunc) *Player {
	return &Player{
		object:  New(id, p),
		Dir:     0,
		Vel:     vec.NewVec2(0, 0),
		Name:    name,
		Size:    size,
		dirFunc: dirFunc,
	}
}

func (p *Player) Update(dt float64, world iface.World) bool {
	p.Dir = p.dirFunc(p, world)
	speed := 10 / float64(p.Size)
	p.Vel = vec.Vec2FromAngle(p.Dir).MulScalar(speed)
	fmt.Println(p.Vel)
	p.pos = p.pos.Add(p.Vel)

	for _, o := range world.GetObjects() {
		if o.Id() == p.id {
			continue
		}
		dist := p.pos.DistanceTo(o.Pos())
		if dist < float64(p.Size) {
			world.RemoveObject(o.Id())
		}
	}

	return true
}
