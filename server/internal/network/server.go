package network

import (
	"log/slog"
	"net"
)

type Server struct {
	clients         map[int]*Client
	newClients      chan *Client
	addr            net.Addr
	receivedPackets chan Packet

	Out chan Packet
}

func NewServer(addr net.Addr) *Server {
	return &Server{
		clients:         make(map[int]*Client, 0),
		newClients:      make(chan *Client, 64),
		addr:            addr,
		receivedPackets: make(chan Packet, 1024),
		Out:             make(chan Packet, 1024),
	}
}

func (s *Server) Start() {
	go s.listen()
	for {
		select {
		case client := <-s.newClients:
			s.assignIdToClient(client)
			s.SendTo(client.Id, &WelcomePacket{clientId: client.Id, playerId: 42})
			go client.Handle(s.receivedPackets)
		case packet := <-s.receivedPackets:
			slog.Debug("received packet", "type", packet)
			s.Out <- packet
		}
	}
}

func (s *Server) listen() {
	addr := net.JoinHostPort(s.addr.String(), "42069")
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		slog.Error("failed to start tcp listener", "error", err)
		return
	}
	slog.Info("listening for connections", "network", s.addr.Network(), "address", s.addr.String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("failed to accept connection", "error", err)
		}
		slog.Info("accepted new tcp connection", "remote", conn.RemoteAddr().String())
		client := NewClient(conn)
		s.newClients <- client
	}
}

func (s *Server) SendTo(clientId int, packet Packet) {
	slog.Info("sending to client", "client id", clientId)
	client := s.clients[clientId]
	n, err := client.tcp.Write(packet.Serialize())
	if err != nil {
		slog.Warn("failed to send packet to client", "client id", clientId, "error", err)
	}
	slog.Debug("wrote bytes to client", "client id", clientId, "byte count", n)
}

func (s *Server) assignIdToClient(client *Client) {
	for id := 0; ; id++ {
		if _, ok := s.clients[id]; ok {
			continue
		}
		client.Id = id
		s.clients[id] = client
		break
	}
	slog.Info("id assigned to client", "id", client.Id, "address", client.tcp.RemoteAddr().String())
}

func (s *Server) GetClients() []Client {
	clients := make([]Client, len(s.clients))
	for _, c := range s.clients {
		clients = append(clients, *c)
	}
	return clients
}
