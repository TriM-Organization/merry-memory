package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

type CreateConstantString struct {
	ConstantString string
}

func (*CreateConstantString) ID() uint16 {
	return IDCreateConstantString
}

func (*CreateConstantString) Name() string {
	return "CreateConstantString"
}

func (c *CreateConstantString) Marshal(io encoding.IO) {
	io.CString(&c.ConstantString)
}
