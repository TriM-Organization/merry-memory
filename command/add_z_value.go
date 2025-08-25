package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddZValue struct{}

func (*AddZValue) ID() uint16 {
	return IDAddZValue
}

func (*AddZValue) Name() string {
	return "AddZValue"
}

func (*AddZValue) Marshal(io encoding.IO) {}
