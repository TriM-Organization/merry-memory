package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt32ZValue struct {
	DeltaZ int32
}

func (*AddInt32ZValue) ID() uint16 {
	return IDAddInt32ZValue
}

func (*AddInt32ZValue) Name() string {
	return "AddInt32ZValue"
}

func (a *AddInt32ZValue) Marshal(io encoding.IO) {
	io.Int32(&a.DeltaZ)
}
