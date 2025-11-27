package main

import (
	"log/slog"
	"net"
	"os"
	"sync"

	"github.com/MatusOllah/slogcolor"
	"github.com/alexrefshauge/agar/server/internal/game/engine"
	"github.com/alexrefshauge/agar/server/internal/network"
	"github.com/alexrefshauge/agar/server/internal/pipe"
)

func main() {
	slogcolor.DefaultOptions.Level = slog.LevelInfo
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))

	engine := engine.New()
	engine.Generate(10)

	addr, err := net.ResolveIPAddr("ip4", "localhost")
	if err != nil {
		panic(err)
	}
	server := network.NewServer(addr)

	pipe := pipe.New(engine, server)
	var wg sync.WaitGroup
	wg.Go(engine.Start)
	wg.Go(server.Start)
	wg.Go(pipe.Start)
	wg.Wait()
}
