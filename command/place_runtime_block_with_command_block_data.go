package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceRuntimeBlockWithCommandBlockData struct {
	BlockRuntimeID   uint16
	CommandBlockData encoding.CommandBlockData
}

func (*PlaceRuntimeBlockWithCommandBlockData) ID() uint16 {
	return IDPlaceRuntimeBlockWithCommandBlockData
}

func (*PlaceRuntimeBlockWithCommandBlockData) Name() string {
	return "PlaceRuntimeBlockWithCommandBlockData"
}

func (p *PlaceRuntimeBlockWithCommandBlockData) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockRuntimeID)
	encoding.Single(io, &p.CommandBlockData)
}
