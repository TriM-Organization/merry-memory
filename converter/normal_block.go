package converter

import (
	"fmt"
	"maps"
	"strings"

	"github.com/TriM-Organization/bedrock-world-operator/block"
	"github.com/TriM-Organization/merry-memory/depends/blocks"
	"github.com/TriM-Organization/merry-memory/utils"
	"github.com/df-mc/worldupgrader/blockupgrader"
	"github.com/df-mc/worldupgrader/itemupgrader"
)

// SetBlock ..
func (c *converter) SetBlock(blockName string, blockStates map[string]any) error {
	if !strings.HasPrefix(blockName, "minecraft:") {
		blockName = "minecraft:" + blockName
	}
	newBlock := blockupgrader.Upgrade(blockupgrader.BlockState{
		Name:       blockName,
		Properties: blockStates,
	})

	blockRuntimeID, found := block.StateToRuntimeID(newBlock.Name, newBlock.Properties)
	if !found {
		blockRuntimeID, _ = block.StateToRuntimeID("minecraft:unknown", map[string]any{})
	}

	err := c.mcworld.SetBlock(c.penPos[0], int16(c.penPos[1]), c.penPos[2], blockRuntimeID)
	if err != nil {
		return fmt.Errorf("SetBlock: %v", err)
	}

	return nil
}

// SetBlockByStatesString ..
func (c *converter) SetBlockByStatesString(blockName string, blockStatesString string) error {
	var blockStates map[string]any

	if !strings.Contains(blockStatesString, "=") {
		blockStates = utils.ParseBlockStatesString(blockStatesString, true)
	} else {
		blockStates = utils.ParseBlockStatesString(blockStatesString, false)
	}

	err := c.SetBlock(blockName, blockStates)
	if err != nil {
		return fmt.Errorf("SetBlockByStatesString: %v", err)
	}
	return nil
}

// SetBlockLegacy ..
func (c *converter) SetBlockLegacy(blockName string, blockData uint16) error {
	copiedStates := make(map[string]any)
	if !strings.HasPrefix(blockName, "minecraft:") {
		blockName = "minecraft:" + blockName
	}

	temp, found := blocks.LegacyBlockToRuntimeID(blockName, blockData)
	if !found {
		newItem := itemupgrader.Upgrade(itemupgrader.ItemMeta{
			Name: blockName,
			Meta: int16(blockData),
		})
		if err := c.SetBlockByStatesString(newItem.Name, "[]"); err != nil {
			return fmt.Errorf("SetBlockLegacy: %v", err)
		}
		return nil
	}

	name, states, _ := blocks.RuntimeIDToState(temp)
	maps.Copy(copiedStates, states)

	err := c.SetBlock(name, copiedStates)
	if err != nil {
		return fmt.Errorf("SetBlockLegacy: %v", err)
	}
	return nil
}

// SetBlockByIndex ..
func (c *converter) SetBlockByIndex(blockNameIndex uint16, blockStatesIndex uint16) error {
	err := c.SetBlockByStatesString(c.constStrings[blockNameIndex], c.constStrings[blockStatesIndex])
	if err != nil {
		return fmt.Errorf("SetBlockByIndex: %v", err)
	}
	return nil
}

// SetBlockLegacyByIndex ..
func (c *converter) SetBlockLegacyByIndex(blockNameIndex uint16, blockData uint16) error {
	err := c.SetBlockLegacy(c.constStrings[blockNameIndex], blockData)
	if err != nil {
		return fmt.Errorf("SetBlockLegacyByIndex: %v", err)
	}
	return nil
}

// SetRuntimeBlock ..
func (c *converter) SetRuntimeBlock(blockRuntimeID uint32) error {
	block := c.runtimeIDPool[blockRuntimeID]
	err := c.SetBlockLegacy(block.Name, block.Data)
	if err != nil {
		return fmt.Errorf("SetRuntimeBlock: %v", err)
	}
	return nil
}
