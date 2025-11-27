package network

import (
	"encoding/binary"
	"math"
)

func uint32BitsFromFloat64(x float64) uint32 {
	return math.Float32bits(float32(x))
}

func (p *WelcomePacket) Serialize() []byte {
	data := make([]byte, headerSize+4+4)
	data[0] = byte(WELCOME)
	binary.BigEndian.PutUint32(data[1:], uint32(p.clientId))
	binary.BigEndian.PutUint32(data[5:], uint32(p.playerId))
	return data
}

const headerSize = 1
const playerNameSize = 16*4
const playerSize, blobSize = 4 + // Id
	4 + // Size
	4 + // Pos X
	4 + // Pos Y
	4 + // Vel X
	4 + // Vel Y
	4 + // Direction
	playerNameSize, // 16 char Name

	4 + // Id
		4 + // Size
		4 + // X
		4 // Y

func (p *StatePacket) Serialize() []byte {
	pCount, bCount := len(p.Players), len(p.Blobs)

	data := make([]byte, headerSize+ // header
		4+4+ // 2 x uint32 for player count and blob count
		pCount*playerSize+
		bCount*blobSize)
	data[0] = byte(STATE)
	binary.BigEndian.PutUint32(data[1+0:], uint32(pCount))
	binary.BigEndian.PutUint32(data[1+4:], uint32(bCount))

	cursor := 1 + 4 + 4
	for _, player := range p.Players {
		binary.BigEndian.PutUint32(data[cursor+0:], uint32(player.Id()))
		binary.BigEndian.PutUint32(data[cursor+4:], uint32(player.Size))
		binary.BigEndian.PutUint32(data[cursor+8:], uint32BitsFromFloat64(player.Pos().X))
		binary.BigEndian.PutUint32(data[cursor+12:], uint32BitsFromFloat64(player.Pos().Y))
		binary.BigEndian.PutUint32(data[cursor+16:], uint32BitsFromFloat64(player.Vel.X))
		binary.BigEndian.PutUint32(data[cursor+20:], uint32BitsFromFloat64(player.Vel.Y))
		binary.BigEndian.PutUint32(data[cursor+24:], uint32BitsFromFloat64(player.Dir))
		copy(data[cursor+28:], []byte(player.Name))
		cursor += playerSize
	}
	for _, blob := range p.Blobs {
		binary.BigEndian.PutUint32(data[cursor+0:], uint32(blob.Id()))
		binary.BigEndian.PutUint32(data[cursor+4:], uint32(blob.Size))
		binary.BigEndian.PutUint32(data[cursor+8:], uint32BitsFromFloat64(blob.Pos().X))
		binary.BigEndian.PutUint32(data[cursor+12:], uint32BitsFromFloat64(blob.Pos().Y))
		cursor += blobSize
	}

	return data
}

func (p *DeltaStatePacket) Serialize() []byte {
	pCount, bCount, unloadCount := len(p.Players), len(p.Blobs), len(p.Unload)

	data := make([]byte, headerSize+ // header
		3*4+ // 3 x uint32 for player count, blob count and unload count
		pCount*playerSize-playerNameSize+
		bCount*blobSize+
		unloadCount*4)

	cursor := 0
	data[0] = byte(DELTA_STATE); cursor++
	binary.BigEndian.PutUint32(data[cursor:], uint32(pCount)); cursor+=4
	binary.BigEndian.PutUint32(data[cursor:], uint32(bCount)); cursor+=4
	binary.BigEndian.PutUint32(data[cursor:], uint32(unloadCount)); cursor+=4

	for _, player := range p.Players {
		binary.BigEndian.PutUint32(data[cursor+0:], uint32(player.Id()))
		binary.BigEndian.PutUint32(data[cursor+4:], uint32(player.Size))
		binary.BigEndian.PutUint32(data[cursor+8:], uint32BitsFromFloat64(player.Pos().X))
		binary.BigEndian.PutUint32(data[cursor+12:], uint32BitsFromFloat64(player.Pos().Y))
		binary.BigEndian.PutUint32(data[cursor+16:], uint32BitsFromFloat64(player.Vel.X))
		binary.BigEndian.PutUint32(data[cursor+20:], uint32BitsFromFloat64(player.Vel.Y))
		binary.BigEndian.PutUint32(data[cursor+24:], uint32BitsFromFloat64(player.Dir))
		cursor += playerSize - playerNameSize
	}
	for _, blob := range p.Blobs {
		binary.BigEndian.PutUint32(data[cursor+0:], uint32(blob.Id()))
		binary.BigEndian.PutUint32(data[cursor+4:], uint32(blob.Size))
		binary.BigEndian.PutUint32(data[cursor+8:], uint32BitsFromFloat64(blob.Pos().X))
		binary.BigEndian.PutUint32(data[cursor+12:], uint32BitsFromFloat64(blob.Pos().Y))
		cursor += blobSize
	}
	for _, id := range p.Unload {
		binary.BigEndian.PutUint32(data[cursor+0:], uint32(id))
		cursor += 4
	}

	return data
}
