package engine

import "github.com/alexrefshauge/agar/server/internal/game/iface"

type Input struct {
	ClientId  int
	Direction float64
}

type Output struct {
	Tick   int
	Update []iface.Object
	Unload []int
}

func NewInput(id int, dir float64) Input {
	return Input{
		ClientId:  id,
		Direction: dir,
	}
}

func NewOutput(tick int, update []iface.Object, unload []int) Output {
	return Output{
		Tick:   tick,
		Update: update,
		Unload: unload,
	}
}
