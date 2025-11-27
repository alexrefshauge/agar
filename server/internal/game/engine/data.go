package engine

import "github.com/alexrefshauge/agar/server/internal/game/iface"

type State map[int]iface.Object

type Input struct {
	ClientId  int
	Direction float64
}

type Output struct {
	Tick   int
	State State
	Updates []int
	Unload []int
}

func NewInput(id int, dir float64) Input {
	return Input{
		ClientId:  id,
		Direction: dir,
	}
}

func NewOutput(tick int, state State, unload []int) Output {
	return Output{
		Tick:   tick,
		State: state,
		Unload: unload,
	}
}
