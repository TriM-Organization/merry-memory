package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceRuntimeBlock struct {
	BlockRuntimeID uint16
}

func (*PlaceRuntimeBlock) ID() uint16 {
	return IDPlaceRuntimeBlock
}

func (*PlaceRuntimeBlock) Name() string {
	return "PlaceRuntimeBlock"
}

func (p *PlaceRuntimeBlock) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockRuntimeID)
}
