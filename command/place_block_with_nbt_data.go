package command

import (
	"github.com/TriM-Organization/merry-memory/protocol/encoding"
	"github.com/TriM-Organization/merry-memory/protocol/nbt"
)

type PlaceBlockWithNBTData struct {
	BlockConstantStringID       uint16
	BlockStatesConstantStringID uint16
	NBTData                     map[string]any
}

func (*PlaceBlockWithNBTData) ID() uint16 {
	return IDPlaceBlockWithNBTData
}

func (*PlaceBlockWithNBTData) Name() string {
	return "PlaceBlockWithNBTData"
}

func (p *PlaceBlockWithNBTData) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockConstantStringID)
	io.Uint16(&p.BlockStatesConstantStringID)
	io.Uint16(&p.BlockStatesConstantStringID) // This is a mistake.
	io.NBT(&p.NBTData, nbt.LittleEndian)
}
