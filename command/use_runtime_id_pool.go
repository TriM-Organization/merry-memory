package command

import (
	"github.com/TriM-Organization/merry-memory/protocol/encoding"
)

type UseRuntimeIDPool struct {
	PoolID uint8
}

func (*UseRuntimeIDPool) ID() uint16 {
	return IDUseRuntimeIDPool
}

func (*UseRuntimeIDPool) Name() string {
	return "UseRuntimeIDPool"
}

func (p *UseRuntimeIDPool) Marshal(io encoding.IO) {
	io.Uint8(&p.PoolID)
}
