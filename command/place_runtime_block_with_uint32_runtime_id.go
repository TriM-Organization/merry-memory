package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceRuntimeBlockWithUint32RuntimeID struct {
	BlockRuntimeID uint32
}

func (*PlaceRuntimeBlockWithUint32RuntimeID) ID() uint16 {
	return IDPlaceRuntimeBlockWithUint32RuntimeID
}

func (*PlaceRuntimeBlockWithUint32RuntimeID) Name() string {
	return "PlaceRuntimeBlockUint32RuntimeID"
}

func (p *PlaceRuntimeBlockWithUint32RuntimeID) Marshal(io encoding.IO) {
	io.Uint32(&p.BlockRuntimeID)
}
