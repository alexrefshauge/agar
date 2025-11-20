package game

import (
	"fmt"
	"log/slog"
	"net"
)

func (g *Game) Run() {
	for {
		//slog.Debug("starting new tick")
		g.newTick()
		//slog.Debug(fmt.Sprintf("new tick at %s\n", g.tickStart.String()))

		//slog.Debug("handling client queue")
		g.handleClientQueue()

		//slog.Debug("running next world step")
		g.world.Step()

		//slog.Debug("broadcasting updates to clients")
		g.broadcastUpdates()

		//slog.Debug("ending tick")
		g.waitTick()
	}
}

func (g *Game) handleClientQueue() {
	for i := 0; i < len(g.clientQueue); i++ {
		conn := <-g.clientQueue
		id := g.addClient(conn)
		slog.Debug("new client", "object id", id)

		packet := fmt.Appendf([]byte{}, "%d;%f\r\n", id, g.world.Size)
		_, err := conn.Write(packet)
		if err != nil {
			g.removeClient(id)
			slog.Error("Failed to add client", "client id", id)
			continue
		}
	}
}

func (g *Game) addClient(conn net.Conn) int {
	p := g.world.AddPlayer()
	id := p.GetId()
	g.clients[id] = Client{
		conn:      conn,
		interests: make(map[int]bool),
	}
	go g.handleClient(id, conn)
	return id
}

func (g *Game) removeClient(id int) {
	client, ok := g.clients[id]
	if !ok {
		slog.Error("Failed to remove client. Client does not exist", "client id", id)
		return
	}

	delete(g.clients, id)
	client.conn.Close()
}
