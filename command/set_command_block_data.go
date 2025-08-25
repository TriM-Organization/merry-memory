package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type SetCommandBlockData struct {
	CommandBlockData encoding.CommandBlockData
}

func (*SetCommandBlockData) ID() uint16 {
	return IDSetCommandBlockData
}

func (*SetCommandBlockData) Name() string {
	return "SetCommandBlockData"
}

func (s *SetCommandBlockData) Marshal(io encoding.IO) {
	encoding.Single(io, &s.CommandBlockData)
}
