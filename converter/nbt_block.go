package converter

import (
	"fmt"
	"maps"
)

// SetNBTBlock ..
func (c *converter) SetNBTBlock(data map[string]any) error {
	nbtData := make(map[string]any)
	maps.Copy(nbtData, data)

	nbtData["x"] = c.penPos[0]
	nbtData["y"] = c.penPos[1]
	nbtData["z"] = c.penPos[2]

	err := c.mcworld.SetBlockNBT(
		c.penPos[0], c.penPos[1], c.penPos[2],
		nbtData,
	)
	if err != nil {
		return fmt.Errorf("SetNBTBlock: %v", err)
	}

	return nil
}
