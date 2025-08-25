package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt8ZValue struct {
	DeltaZ int8
}

func (*AddInt8ZValue) ID() uint16 {
	return IDAddInt8ZValue
}

func (*AddInt8ZValue) Name() string {
	return "AddInt8ZValue"
}

func (a *AddInt8ZValue) Marshal(io encoding.IO) {
	io.Int8(&a.DeltaZ)
}
