package engine

import (
	"log/slog"
	"time"

	"github.com/alexrefshauge/agar/server/internal/game/iface"
	"github.com/alexrefshauge/agar/server/internal/game/world"
)

type Engine struct {
	world    iface.World
	ticker   *time.Ticker
	lastTick time.Time
	In       chan Input
	Out      chan Output
}

func New() *Engine {
	return &Engine{
		world: world.New(1000),
		In:    make(chan Input, 1024),
		Out:   make(chan Output, 1024),
	}
}

func (e *Engine) Start() {
	duration := 100 * time.Millisecond
	slog.Info("starting engine", "tick duration", duration)
	e.ticker = time.NewTicker(duration)
	for tick := 0; ; tick++ {
		t := <-e.ticker.C
		e.Step(t.Sub(e.lastTick).Seconds(), tick)
		e.lastTick = t
	}
}

func (e *Engine) Step(dt float64, tick int) {
	updates, removals := e.updateObjects(dt)
	e.Out <- NewOutput(tick, updates, removals)
	slog.Debug("objects updated", "updates", len(updates), "removals", len(removals))
}

// updateObjects updates all game objects with delta time,
// and returns an []int with the ids of each object that changed during update
func (e *Engine) updateObjects(dt float64) ([]iface.Object, []int) {
	e.world.Cleanup()
	updates := make([]iface.Object, 0)
	for _, o := range e.world.GetObjects() {
		if o.Update(dt, e.world) {
			updates = append(updates, o)
		}
		updates = append(updates, o)
	}
	return updates, e.world.GetRemovals()
}
