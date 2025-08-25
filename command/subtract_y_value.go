package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type SubtractYValue struct{}

func (*SubtractYValue) ID() uint16 {
	return IDSubtractYValue
}

func (*SubtractYValue) Name() string {
	return "SubtractYValue"
}

func (*SubtractYValue) Marshal(io encoding.IO) {}
