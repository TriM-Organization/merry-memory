package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID struct {
	BlockRuntimeID   uint32
	CommandBlockData encoding.CommandBlockData
}

func (*PlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID) ID() uint16 {
	return IDPlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID
}

func (*PlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID) Name() string {
	return "PlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID"
}

func (p *PlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID) Marshal(io encoding.IO) {
	io.Uint32(&p.BlockRuntimeID)
	encoding.Single(io, &p.CommandBlockData)
}
