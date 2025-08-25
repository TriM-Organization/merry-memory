package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceBlockWithBlockStates struct {
	BlockConstantStringID       uint16
	BlockStatesConstantStringID uint16
}

func (*PlaceBlockWithBlockStates) ID() uint16 {
	return IDPlaceBlockWithBlockStates
}

func (*PlaceBlockWithBlockStates) Name() string {
	return "PlaceBlockWithBlockStates"
}

func (p *PlaceBlockWithBlockStates) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockConstantStringID)
	io.Uint16(&p.BlockStatesConstantStringID)
}
