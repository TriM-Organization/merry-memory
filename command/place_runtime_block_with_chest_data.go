package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceRuntimeBlockWithChestData struct {
	BlockRuntimeID uint16
	ChestSlots     []encoding.ChestSlot
}

func (*PlaceRuntimeBlockWithChestData) ID() uint16 {
	return IDPlaceRuntimeBlockWithChestData
}

func (*PlaceRuntimeBlockWithChestData) Name() string {
	return "PlaceRuntimeBlockWithChestData"
}

func (p *PlaceRuntimeBlockWithChestData) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockRuntimeID)
	encoding.SliceUint8Length(io, &p.ChestSlots)
}
