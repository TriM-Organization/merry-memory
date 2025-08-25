package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt32YValue struct {
	DeltaY int32
}

func (*AddInt32YValue) ID() uint16 {
	return IDAddInt32YValue
}

func (*AddInt32YValue) Name() string {
	return "AddInt32YValue"
}

func (a *AddInt32YValue) Marshal(io encoding.IO) {
	io.Int32(&a.DeltaY)
}
