package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type PlaceCommandBlockWithCommandBlockData struct {
	BlockData        uint16
	CommandBlockData encoding.CommandBlockData
}

func (*PlaceCommandBlockWithCommandBlockData) ID() uint16 {
	return IDPlaceCommandBlockWithCommandBlockData
}

func (*PlaceCommandBlockWithCommandBlockData) Name() string {
	return "PlaceCommandBlockWithCommandBlockData"
}

func (p *PlaceCommandBlockWithCommandBlockData) Marshal(io encoding.IO) {
	io.Uint16(&p.BlockData)
	encoding.Single(io, &p.CommandBlockData)
}
