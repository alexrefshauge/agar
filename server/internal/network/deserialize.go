package network

import "encoding/binary"

func Deserialize(data []byte) Packet {
	packetType := uint8(data[0])
	payload := data[1:]
	switch packetType {
	case STATE:
		return nil
	case DELTA_STATE:
		return nil
	case PLAYER_INPUT:
		return deserializePlayerInput(payload)
	default:
		return nil
	}
}

func deserializePlayerInput(payload []byte) *PlayerInputPacket {
	clientId := binary.BigEndian.Uint32(payload[0:3])
	direction := binary.BigEndian.Uint32(payload[4:])
	return &PlayerInputPacket{
		ClientId:  int(clientId),
		Direction: float64(direction),
	}
}
