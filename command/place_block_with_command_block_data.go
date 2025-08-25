package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceBlockWithCommandBlockData struct {
	BlockConstantStringID uint16
	BlockData             uint16
	CommandBlockData      encoding.CommandBlockData
}

func (*PlaceBlockWithCommandBlockData) ID() uint16 {
	return IDPlaceBlockWithCommandBlockData
}

func (*PlaceBlockWithCommandBlockData) Name() string {
	return "PlaceBlockWithCommandBlockData"
}

func (p *PlaceBlockWithCommandBlockData) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockConstantStringID)
	io.Uint16(&p.BlockData)
	encoding.Single(io, &p.CommandBlockData)
}
