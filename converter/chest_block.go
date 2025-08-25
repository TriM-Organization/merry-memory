package converter

import (
	"fmt"

	"github.com/TriM-Organization/merry-memory/protocol/encoding"
)

// SetChestBlock ..
func (c *converter) SetChestBlock(data []encoding.ChestSlot) error {
	items := make([]any, 0)
	for _, value := range data {
		items = append(items, map[string]any{
			"Count":  byte(value.Count),
			"Damage": int16(value.Damage),
			"Name":   value.Name,
			"Slot":   byte(value.Slot),
		})
	}

	err := c.mcworld.SetBlockNBT(
		c.penPos[0],
		c.penPos[1],
		c.penPos[2],
		map[string]any{
			"Items": items,
			"x":     c.penPos[0],
			"y":     c.penPos[1],
			"z":     c.penPos[2],
		},
	)
	if err != nil {
		return fmt.Errorf("SetCommandBlock: %v", err)
	}

	return nil
}
