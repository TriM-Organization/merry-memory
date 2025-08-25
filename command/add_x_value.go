package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddXValue struct{}

func (*AddXValue) ID() uint16 {
	return IDAddXValue
}

func (*AddXValue) Name() string {
	return "AddXValue"
}

func (*AddXValue) Marshal(io encoding.IO) {}
