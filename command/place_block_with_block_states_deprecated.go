package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceBlockWithBlockStatesDeprecated struct {
	BlockConstantStringID uint16
	BlockStatesString     string
}

func (*PlaceBlockWithBlockStatesDeprecated) ID() uint16 {
	return IDPlaceBlockWithBlockStatesDeprecated
}

func (*PlaceBlockWithBlockStatesDeprecated) Name() string {
	return "PlaceBlockWithBlockStatesDeprecated"
}

func (p *PlaceBlockWithBlockStatesDeprecated) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockConstantStringID)
	io.CString(&p.BlockStatesString)
}
