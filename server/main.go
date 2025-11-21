package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/alexrefshauge/agar/server/game"
	"github.com/alexrefshauge/agar/server/game/world"
)

var HOST = "localhost"
var PORT = 42069
var gameWorld *world.World

func init() {
	gameWorld = world.NewWorld()
	gameWorld.Generate(100)
}

func main() {
	if len(os.Args) >= 3 {
		HOST = os.Args[1]
		p, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			panic(err)
		}
		PORT = int(p)
	}

	slog.SetLogLoggerLevel(slog.LevelDebug)
	address := fmt.Sprintf("%s:%d", HOST, PORT)
	fmt.Printf("starting server on %s\n", address)

	g := game.NewGameWithWorld(*gameWorld)
	go g.Listen(address)
	g.Run()
}
