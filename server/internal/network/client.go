package network

import (
	"log/slog"
	"net"
)

type Client struct {
	Id       int
	PlayerId int
	Active   bool

	tcp           net.Conn
	udp           net.Conn
	broadcastChan chan Packet
	interest      map[int]bool
}

func NewClient(conn net.Conn) *Client {
	slog.Info("establishing new client connection")
	//welcome := WelcomePacket{clientId: -1, playerId: -1}
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero})
	if err != nil {
		slog.Error("failed to start udp connection", "error", err)
		return nil
	}

	return &Client{
		Id:            -1,
		PlayerId:      -1,
		Active:        true,
		interest:      make(map[int]bool),
		broadcastChan: make(chan Packet),
		tcp:           conn,
		udp:           udpConn,
	}
}

func (c *Client) Handle(receive chan Packet) {
	dataBuffer := make([]byte, 4096)
	defer func() { c.udp.Close(); c.tcp.Close() }()

	for {
		_, err := c.tcp.Read(dataBuffer)
		typedata := dataBuffer[0]
		if err != nil {
			c.tcp.Close()
			slog.Warn("failed to read from client", "error", err)
			break
		}
		packetType := uint8(typedata)
		switch packetType {
		case PLAYER_INPUT:
			deserializePlayerInput(dataBuffer) //TODO: read more safely
		}
	}

	slog.Info("client disconnected", "client id", c.Id)
	//	go func() {
	//		time.Sleep(30 * time.Second)
	//		c.Active = false
	//	}()
}
