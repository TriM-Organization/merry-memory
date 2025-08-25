package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type Terminate struct{}

func (*Terminate) ID() uint16 {
	return IDTerminate
}

func (*Terminate) Name() string {
	return "Terminate"
}

func (*Terminate) Marshal(io encoding.IO) {}
