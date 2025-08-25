package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt16YValue struct {
	DeltaY int16
}

func (*AddInt16YValue) ID() uint16 {
	return IDAddInt16YValue
}

func (*AddInt16YValue) Name() string {
	return "AddInt16YValue"
}

func (a *AddInt16YValue) Marshal(io encoding.IO) {
	io.Int16(&a.DeltaY)
}
