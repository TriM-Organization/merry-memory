package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type SubtractZValue struct{}

func (*SubtractZValue) ID() uint16 {
	return IDSubtractZValue
}

func (*SubtractZValue) Name() string {
	return "SubtractZValue"
}

func (*SubtractZValue) Marshal(io encoding.IO) {}
