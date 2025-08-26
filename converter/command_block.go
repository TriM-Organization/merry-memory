package converter

import (
	"fmt"

	"github.com/TriM-Organization/bedrock-world-operator/block"
	"github.com/TriM-Organization/merry-memory/protocol/encoding"
)

// SetCommandBlock ..
func (c *converter) SetCommandBlock(data encoding.CommandBlockData) error {
	nbtMap := map[string]any{
		"Command":    data.Command,
		"CustomName": data.CustomName,
		"LastOutput": data.LastOutput,
		"TickDelay":  data.TickDelay,
		"x":          c.penPos[0],
		"y":          c.penPos[1],
		"z":          c.penPos[2],
		"id":         "CommandBlock",
	}

	if data.ExecuteOnFirstTick {
		nbtMap["ExecuteOnFirstTick"] = byte(1)
	} else {
		nbtMap["ExecuteOnFirstTick"] = byte(0)
	}
	if data.TrackOutput {
		nbtMap["TrackOutput"] = byte(1)
	} else {
		nbtMap["TrackOutput"] = byte(0)
	}
	if data.Conditional {
		nbtMap["conditionalMode"] = byte(1)
	} else {
		nbtMap["conditionalMode"] = byte(0)
	}
	if !data.NeedsRedstone {
		nbtMap["auto"] = byte(1)
	} else {
		nbtMap["auto"] = byte(0)
	}

	err := c.mcworld.SetBlockNBT(c.penPos[0], c.penPos[1], c.penPos[2], nbtMap)
	if err != nil {
		return fmt.Errorf("SetCommandBlock: %v", err)
	}

	blockRuntimeID, err := c.mcworld.LoadBlock(c.penPos[0], int16(c.penPos[1]), c.penPos[2])
	if err != nil {
		return fmt.Errorf("SetCommandBlock: %v", err)
	}
	_, blockStates, found := block.RuntimeIDToState(blockRuntimeID)
	if !found {
		return nil
	}

	switch data.Mode {
	case 0:
		err = c.SetBlock("minecraft:command_block", blockStates)
	case 1:
		err = c.SetBlock("minecraft:repeating_command_block", blockStates)
	case 2:
		err = c.SetBlock("minecraft:chain_command_block", blockStates)
	}
	if err != nil {
		return fmt.Errorf("SetCommandBlock: %v", err)
	}

	return nil
}
