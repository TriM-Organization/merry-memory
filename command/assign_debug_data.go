package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type AssignDebugData struct {
	Data []byte
}

func (*AssignDebugData) ID() uint16 {
	return IDAssignDebugData
}

func (*AssignDebugData) Name() string {
	return "AssignDebugData"
}

func (a *AssignDebugData) Marshal(io encoding.IO) {
	encoding.FuncSliceUint32Length(io, &a.Data, io.Uint8)
}
