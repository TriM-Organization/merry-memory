package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddYValue struct{}

func (*AddYValue) ID() uint16 {
	return IDAddYValue
}

func (*AddYValue) Name() string {
	return "AddYValue"
}

func (*AddYValue) Marshal(io encoding.IO) {}
