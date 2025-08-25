package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type SubtractXValue struct{}

func (*SubtractXValue) ID() uint16 {
	return IDSubtractXValue
}

func (*SubtractXValue) Name() string {
	return "SubtractXValue"
}

func (*SubtractXValue) Marshal(io encoding.IO) {}
