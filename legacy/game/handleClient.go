package game

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"strconv"

	"github.com/alexrefshauge/agar/server/game/object"
)

func (g *Game) handleClient(id int, conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	partial := ""

	for {
		data, prefix, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Warn("reading from closed socket", "client", id)
				break
			}
			slog.Error("failed to read from socket", "error", err)
			g.removeClient(id)
			return
		}
		if prefix {
			partial = fmt.Sprintf("%s%s", partial, data)
			continue
		}
		packet := fmt.Sprintf("%s%s", partial, data)
		partial = ""

		direction, err := strconv.ParseFloat(string(packet), 64)
		o, ok := g.world.IdMap[id]
		if !ok {
			//TODO: handle gracefully
			fmt.Printf("Player with client id %d does not exist. Closing connection", id)
			conn.Close()
			return
		}

		switch obj := o.(type) {
		case *object.Player:
			obj.Direction = float32(direction)
			break
		default:
			fmt.Println("GameObject id and client id mismatch. GameObject is not a Player")
		}
	}
}
