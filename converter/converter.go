package converter

import (
	"context"
	"fmt"

	"github.com/TriM-Organization/merry-memory/command"
	"github.com/TriM-Organization/merry-memory/depends/runtime_id_pool"
	"github.com/TriM-Organization/merry-memory/utils"
)

// converter ..
type converter struct {
	mcworld       *utils.MCWorld
	constStrings  []string
	runtimeIDPool []*runtime_id_pool.ConstBlock
	penPos        utils.BlockPos
}

// ConvertBDXToMCWorld ..
func ConvertBDXToMCWorld(bdxPath string, mcworldPath string) error {
	_, reader, err := utils.ReadBDXFileInfo(bdxPath)
	if err != nil {
		return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	mcworld, err := utils.NewMCWorld(mcworldPath, ctx)
	if err != nil {
		return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
	}

	converter := converter{
		mcworld:      mcworld,
		constStrings: nil,
		penPos:       utils.BlockPos{0, -64, 0},
	}

	for {
		cmd, err := utils.ReadBDXCommand(reader)
		if err != nil {
			return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
		}

		switch c := cmd.(type) {
		case *command.CreateConstantString:
			converter.constStrings = append(converter.constStrings, c.ConstantString)
		case *command.PlaceBlockWithBlockStates:
			err = converter.SetBlockByIndex(c.BlockConstantStringID, c.BlockStatesConstantStringID)
		case *command.AddInt16ZValue0:
			converter.penPos[2] += int32(c.DeltaZ)
		case *command.PlaceBlock:
			err = converter.SetBlockLegacyByIndex(c.BlockConstantStringID, c.BlockData)
		case *command.AddZValue0:
			converter.penPos[2] += 1
		case *command.NoOperation:
		case *command.AddInt32ZValue0:
			converter.penPos[2] += c.DeltaZ
		case *command.PlaceBlockWithBlockStatesDeprecated:
			err = converter.SetBlockByStatesString(converter.constStrings[c.BlockConstantStringID], c.BlockStatesString)
		case *command.AddXValue:
			converter.penPos[0] += 1
		case *command.SubtractXValue:
			converter.penPos[0] -= 1
		case *command.AddYValue:
			converter.penPos[1] += 1
		case *command.SubtractYValue:
			converter.penPos[1] -= 1
		case *command.AddZValue:
			converter.penPos[2] += 1
		case *command.SubtractZValue:
			converter.penPos[2] -= 1
		case *command.AddInt16XValue:
			converter.penPos[0] += int32(c.DeltaX)
		case *command.AddInt32XValue:
			converter.penPos[0] += c.DeltaX
		case *command.AddInt16YValue:
			converter.penPos[1] += int32(c.DeltaY)
		case *command.AddInt32YValue:
			converter.penPos[1] += c.DeltaY
		case *command.AddInt16ZValue:
			converter.penPos[2] += int32(c.DeltaZ)
		case *command.AddInt32ZValue:
			converter.penPos[2] += c.DeltaZ
		case *command.SetCommandBlockData:
			err = converter.SetCommandBlock(c.CommandBlockData)
		case *command.PlaceBlockWithCommandBlockData:
			if err = converter.SetBlockLegacyByIndex(c.BlockConstantStringID, c.BlockData); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
			if err = converter.SetCommandBlock(c.CommandBlockData); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
		case *command.AddInt8XValue:
			converter.penPos[0] += int32(c.DeltaX)
		case *command.AddInt8YValue:
			converter.penPos[1] += int32(c.DeltaY)
		case *command.AddInt8ZValue:
			converter.penPos[2] += int32(c.DeltaZ)
		case *command.UseRuntimeIDPool:
			switch c.PoolID {
			case 117:
				converter.runtimeIDPool = runtime_id_pool.RuntimeIdArray_117
			case 118:
				converter.runtimeIDPool = runtime_id_pool.RuntimeIdArray_2_1_10
			default:
				err = fmt.Errorf(
					"ConvertBDXToMCWorld: This file is using an unknown runtime id pool %d, we're unable to resolve it",
					c.PoolID,
				)
			}
		case *command.PlaceRuntimeBlock:
			err = converter.SetRuntimeBlock(uint32(c.BlockRuntimeID))
		case *command.PlaceRuntimeBlockWithUint32RuntimeID:
			err = converter.SetRuntimeBlock(c.BlockRuntimeID)
		case *command.PlaceRuntimeBlockWithCommandBlockData:
			if err = converter.SetRuntimeBlock(uint32(c.BlockRuntimeID)); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
			if err = converter.SetCommandBlock(c.CommandBlockData); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
		case *command.PlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID:
			if err = converter.SetRuntimeBlock(c.BlockRuntimeID); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
			if err = converter.SetCommandBlock(c.CommandBlockData); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
		case *command.PlaceCommandBlockWithCommandBlockData:
			if err = converter.SetBlockLegacy("minecraft:command_block", c.BlockData); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
			if err = converter.SetCommandBlock(c.CommandBlockData); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
		case *command.PlaceRuntimeBlockWithChestData:
			if err = converter.SetRuntimeBlock(uint32(c.BlockRuntimeID)); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
			if err = converter.SetChestBlock(c.ChestSlots); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
		case *command.PlaceRuntimeBlockWithChestDataAndUint32RuntimeID:
			if err = converter.SetRuntimeBlock(c.BlockRuntimeID); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
			if err = converter.SetChestBlock(c.ChestSlots); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
		case *command.AssignDebugData:
			// Can support, but need ask Liliya233 and CMA2401PT
		case *command.PlaceBlockWithChestData:
			if err = converter.SetBlockLegacyByIndex(c.BlockConstantStringID, c.BlockData); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
			if err = converter.SetChestBlock(c.ChestSlots); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
		case *command.PlaceBlockWithNBTData:
			err = converter.SetNBTBlock(c.NBTData)
		case *command.Terminate:
			if err = converter.mcworld.Close(); err != nil {
				return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
			}
			return nil
		}

		if err != nil {
			return fmt.Errorf("ConvertBDXToMCWorld: %v", err)
		}
	}
}
