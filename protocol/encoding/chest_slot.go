package encoding

// ChestSlot ..
type ChestSlot struct {
	Name   string
	Count  uint8
	Damage uint16
	Slot   uint8
}

func (c *ChestSlot) Marshal(io IO) {
	io.CString(&c.Name)
	io.Uint8(&c.Count)
	io.Uint16(&c.Damage)
	io.Uint8(&c.Slot)
}
