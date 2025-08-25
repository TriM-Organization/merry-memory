package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt16ZValue struct {
	DeltaZ int16
}

func (*AddInt16ZValue) ID() uint16 {
	return IDAddInt16ZValue
}

func (*AddInt16ZValue) Name() string {
	return "AddInt16ZValue"
}

func (a *AddInt16ZValue) Marshal(io encoding.IO) {
	io.Int16(&a.DeltaZ)
}
