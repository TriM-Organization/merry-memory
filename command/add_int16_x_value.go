package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt16XValue struct {
	DeltaX int16
}

func (*AddInt16XValue) ID() uint16 {
	return IDAddInt16XValue
}

func (*AddInt16XValue) Name() string {
	return "AddInt16XValue"
}

func (a *AddInt16XValue) Marshal(io encoding.IO) {
	io.Int16(&a.DeltaX)
}
