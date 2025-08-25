package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceBlockWithChestData struct {
	BlockConstantStringID uint16
	BlockData             uint16
	ChestSlots            []encoding.ChestSlot
}

func (*PlaceBlockWithChestData) ID() uint16 {
	return IDPlaceBlockWithChestData
}

func (*PlaceBlockWithChestData) Name() string {
	return "PlaceBlockWithChestData"
}

func (p *PlaceBlockWithChestData) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockConstantStringID)
	io.Uint16(&p.BlockData)
	encoding.SliceUint8Length(io, &p.ChestSlots)
}
