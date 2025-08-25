package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AddInt32ZValue0 struct {
	DeltaZ int32
}

func (*AddInt32ZValue0) ID() uint16 {
	return IDAddInt32ZValue0
}

func (*AddInt32ZValue0) Name() string {
	return "AddInt32ZValue0"
}

func (a *AddInt32ZValue0) Marshal(io encoding.IO) {
	io.Int32(&a.DeltaZ)
}
