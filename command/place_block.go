package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceBlock struct {
	BlockConstantStringID uint16
	BlockData             uint16
}

func (*PlaceBlock) ID() uint16 {
	return IDPlaceBlock
}

func (*PlaceBlock) Name() string {
	return "PlaceBlock"
}

func (p *PlaceBlock) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockConstantStringID)
	io.Uint16(&p.BlockData)
}
