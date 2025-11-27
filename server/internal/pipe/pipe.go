package pipe

import (
	"github.com/alexrefshauge/agar/server/internal/game/engine"
	"github.com/alexrefshauge/agar/server/internal/game/object"
	"github.com/alexrefshauge/agar/server/internal/network"
)

type Pipe struct {
	engine *engine.Engine
	server *network.Server
}

func New(e *engine.Engine, s *network.Server) *Pipe {
	return &Pipe{engine: e, server: s}
}

func (pipe *Pipe) Start() {
	for {
		select {

		case packet := <-pipe.server.Out: //Received packet from server -> put into engine
			pipe.processServerPacket(packet)

		case out := <-pipe.engine.Out: // Data came out of the engine -> broadcast state
			pipe.processEngineOutput(out)
		}
	}
}

func (pipe *Pipe) processServerPacket(p network.Packet) {
	switch p := p.(type) {
	case *network.PlayerInputPacket:
		pipe.engine.In <- engine.NewInput(p.ClientId, p.Direction)
	}
}

func (pipe *Pipe) processEngineOutput(output engine.Output) {
	packet := &network.StatePacket{
		Players: make([]*object.Player, 0),
		Blobs:   make([]*object.Blob, 0),
	}
	for _, o := range output.Update {
		if player, ok := o.(*object.Player); ok {
			packet.Players = append(packet.Players, player)
		}
		if blob, ok := o.(*object.Blob); ok {
			packet.Blobs = append(packet.Blobs, blob)
		}
	}
	for _, c := range pipe.server.GetClients() {
		pipe.server.SendTo(c.Id, packet)
	}
}
