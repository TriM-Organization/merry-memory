package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceRuntimeBlockWithChestDataAndUint32RuntimeID struct {
	BlockRuntimeID uint32
	ChestSlots     []encoding.ChestSlot
}

func (*PlaceRuntimeBlockWithChestDataAndUint32RuntimeID) ID() uint16 {
	return IDPlaceRuntimeBlockWithChestDataAndUint32RuntimeID
}

func (*PlaceRuntimeBlockWithChestDataAndUint32RuntimeID) Name() string {
	return "PlaceRuntimeBlockWithChestDataAndUint32RuntimeID"
}

func (p *PlaceRuntimeBlockWithChestDataAndUint32RuntimeID) Marshal(io encoding.IO) {
	io.Uint32(&p.BlockRuntimeID)
	encoding.SliceUint8Length(io, &p.ChestSlots)
}
