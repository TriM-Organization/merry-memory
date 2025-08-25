package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt32XValue struct {
	DeltaX int32
}

func (*AddInt32XValue) ID() uint16 {
	return IDAddInt32XValue
}

func (*AddInt32XValue) Name() string {
	return "AddInt32XValue"
}

func (a *AddInt32XValue) Marshal(io encoding.IO) {
	io.Int32(&a.DeltaX)
}
