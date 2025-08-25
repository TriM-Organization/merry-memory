package encoding

// CommandBlockData ..
type CommandBlockData struct {
	Mode               uint32
	Command            string
	CustomName         string
	LastOutput         string
	TickDelay          int32
	ExecuteOnFirstTick bool
	TrackOutput        bool
	Conditional        bool
	NeedsRedstone      bool
}

func (c *CommandBlockData) Marshal(io IO) {
	io.Uint32(&c.Mode)
	io.CString(&c.Command)
	io.CString(&c.CustomName)
	io.CString(&c.LastOutput)
	io.Int32(&c.TickDelay)
	io.Bool(&c.ExecuteOnFirstTick)
	io.Bool(&c.TrackOutput)
	io.Bool(&c.Conditional)
	io.Bool(&c.NeedsRedstone)
}
