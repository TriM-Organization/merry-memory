package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt16ZValue0 struct {
	DeltaZ int16
}

func (*AddInt16ZValue0) ID() uint16 {
	return IDAddInt16ZValue0
}

func (*AddInt16ZValue0) Name() string {
	return "AddInt16ZValue0"
}

func (a *AddInt16ZValue0) Marshal(io encoding.IO) {
	io.Int16(&a.DeltaZ)
}
