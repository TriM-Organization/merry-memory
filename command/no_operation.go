package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type NoOperation struct{}

func (*NoOperation) ID() uint16 {
	return IDNOP
}

func (*NoOperation) Name() string {
	return "NOP"
}

func (*NoOperation) Marshal(io encoding.IO) {}
