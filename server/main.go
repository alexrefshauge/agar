package main

import (
	"fmt"

	"github.com/alexrefshauge/agar/server/game"
	"github.com/alexrefshauge/agar/server/game/world"
)

var HOST = "localhost"
var PORT = 42069
var gameWorld *world.World

func init() {
	gameWorld = world.NewWorld()
	gameWorld.Generate(25)
}

func main() {
	address := fmt.Sprintf("%s:%d", HOST, PORT)
	fmt.Printf("starting server on %s\n", address)

	g := game.NewGameWithWorld(*gameWorld)
	go g.Run()

	g.Listen(address)
}
