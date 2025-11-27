package pipe

import (
	"github.com/alexrefshauge/agar/server/internal/game/engine"
	"github.com/alexrefshauge/agar/server/internal/game/object"
	"github.com/alexrefshauge/agar/server/internal/network"
)

type clientInterests map[int]bool

type Pipe struct {
	engine *engine.Engine
	server *network.Server
	interests map[int]clientInterests
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
	players := make([]*object.Player, 0)
	blobs := make([]*object.Blob, 0)

	for _, o := range output.State {
		if player, ok := o.(*object.Player); ok {
			players = append(players, player)
		}
		if blob, ok := o.(*object.Blob); ok {
			blobs = append(blobs, blob)
		}
	}

	for _, c := range pipe.server.GetClients() {
		if !c.Active {
			continue
		}

		pipe.sendStateToClient(c, players, blobs, output.State, output.Unload)
	}
}

func (pipe *Pipe) sendStateToClient(
		client network.Client, 
		players []*object.Player, 
		blobs []*object.Blob,
		state engine.State,
		unloads []int) {
	var packet network.Packet
	pipe.updateClientInterests(client, state)

	if _, hasInterests := pipe.interests[client.Id]; hasInterests {
		packet = &network.DeltaStatePacket{
			Players: players,
			Blobs:   blobs,
			Unload: unloads,
		}
	} else {
		packet = &network.StatePacket{
			Players: players,
			Blobs:   blobs,
		}
	}


	pipe.server.SendTo(client.Id, packet)
}

func (pipe *Pipe) updateClientInterests(client network.Client, state engine.State) {
	
}