package converter

import (
	"fmt"
	"strings"

	"github.com/TriM-Organization/bedrock-world-operator/block"
	"github.com/TriM-Organization/merry-memory/protocol/encoding"
)

// SetContainerData ..
func (c *converter) SetContainerData(data []encoding.ChestSlot) error {
	var blockNBTID string

	items := make([]any, 0)
	nbtMap := map[string]any{
		"x": c.penPos[0],
		"y": c.penPos[1],
		"z": c.penPos[2],
	}
	for _, value := range data {
		items = append(items, map[string]any{
			"Count":  byte(value.Count),
			"Damage": int16(value.Damage),
			"Name":   value.Name,
			"Slot":   byte(value.Slot),
		})
	}

	blockRuntimeID, err := c.mcworld.LoadBlock(c.penPos[0], int16(c.penPos[1]), c.penPos[2])
	if err != nil {
		return fmt.Errorf("SetCommandBlock: %v", err)
	}
	blockName, _, _ := block.RuntimeIDToState(blockRuntimeID)

	switch blockName {
	case "minecraft:blast_furnace", "minecraft:lit_blast_furnace":
		blockNBTID = "BlastFurnace"
	case "minecraft:furnace", "minecraft:lit_furnace":
		blockNBTID = "Furnace"
	case "minecraft:smoker", "minecraft:lit_smoker":
		blockNBTID = "Smoker"
	case "minecraft:chest", "minecraft:trapped_chest":
		blockNBTID = "Chest"
	case "minecraft:hopper":
		blockNBTID = "Hopper"
	case "minecraft:dispenser":
		blockNBTID = "Dispenser"
	case "minecraft:dropper":
		blockNBTID = "Dropper"
	case "minecraft:barrel":
		blockNBTID = "Barrel"
	case "minecraft:crafter":
		blockNBTID = "Crafter"
	}
	if strings.Contains(blockName, "shulker_box") {
		blockNBTID = "ShulkerBox"
		nbtMap["facing"] = byte(1)
	}

	nbtMap["Items"] = items
	nbtMap["id"] = blockNBTID
	if err = c.mcworld.SetBlockNBT(c.penPos[0], c.penPos[1], c.penPos[2], nbtMap); err != nil {
		return fmt.Errorf("SetCommandBlock: %v", err)
	}
	return nil
}
