package game

import (
	"log/slog"
	"time"
)

var tickRate = 100 * time.Millisecond

func (g *Game) newTick() {
	g.tickStart = time.Now()
}

func (g *Game) waitTick() {
	tickDuraction := time.Since(g.tickStart)
	if tickDuraction > tickRate {
		slog.Warn("slow tick", "duration", tickDuraction.Milliseconds())
		return
	}

	time.Sleep(tickRate - tickDuraction)
}
