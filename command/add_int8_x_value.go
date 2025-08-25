package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt8XValue struct {
	DeltaX int8
}

func (*AddInt8XValue) ID() uint16 {
	return IDAddInt8XValue
}

func (*AddInt8XValue) Name() string {
	return "AddInt8XValue"
}

func (a *AddInt8XValue) Marshal(io encoding.IO) {
	io.Int8(&a.DeltaX)
}
