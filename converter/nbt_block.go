package converter

import (
	"fmt"
	"maps"
)

// SetNBTBlock ..
func (c *converter) SetNBTBlock(data map[string]any) error {
	nbtData := make(map[string]any)
	maps.Copy(nbtData, data)

	nbtData = c.processPairChest(nbtData)
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

// processPairChest ..
func (c *converter) processPairChest(data map[string]any) map[string]any {
	blockNBTID, _ := data["id"].(string)
	if blockNBTID != "Chest" {
		return data
	}

	_, ok := data["pairlead"].(byte)
	if !ok {
		return data
	}

	posX, ok := data["x"].(int32)
	if !ok {
		return data
	}
	posZ, ok := data["z"].(int32)
	if !ok {
		return data
	}

	pairX, ok := data["pairx"].(int32)
	if !ok {
		return data
	}
	pairZ, ok := data["pairz"].(int32)
	if !ok {
		return data
	}

	data["pairx"] = pairX - posX + c.penPos[0]
	data["pairz"] = pairZ - posZ + c.penPos[2]

	return data
}
