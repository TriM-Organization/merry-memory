package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt8YValue struct {
	DeltaY int8
}

func (*AddInt8YValue) ID() uint16 {
	return IDAddInt8YValue
}

func (*AddInt8YValue) Name() string {
	return "AddInt8YValue"
}

func (a *AddInt8YValue) Marshal(io encoding.IO) {
	io.Int8(&a.DeltaY)
}
