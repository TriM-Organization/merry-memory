package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddZValue0 struct{}

func (*AddZValue0) ID() uint16 {
	return IDAddZValue0
}

func (*AddZValue0) Name() string {
	return "AddZValue0"
}

func (*AddZValue0) Marshal(io encoding.IO) {}
