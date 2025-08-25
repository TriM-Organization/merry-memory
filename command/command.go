package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

const (
	IDCreateConstantString = iota + 1
	_
	_
	_
	IDPlaceBlockWithBlockStates
	IDAddInt16ZValue0
	IDPlaceBlock
	IDAddZValue0
	IDNOP
	_
	_
	IDAddInt32ZValue0
	IDPlaceBlockWithBlockStatesDeprecated
	IDAddXValue
	IDSubtractXValue
	IDAddYValue
	IDSubtractYValue
	IDAddZValue
	IDSubtractZValue
	IDAddInt16XValue
	IDAddInt32XValue
	IDAddInt16YValue
	IDAddInt32YValue
	IDAddInt16ZValue
	IDAddInt32ZValue
	IDSetCommandBlockData
	IDPlaceBlockWithCommandBlockData
	IDAddInt8XValue
	IDAddInt8YValue
	IDAddInt8ZValue
	IDUseRuntimeIDPool
	IDPlaceRuntimeBlock
	IDPlaceRuntimeBlockWithUint32RuntimeID
	IDPlaceRuntimeBlockWithCommandBlockData
	IDPlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID
	IDPlaceCommandBlockWithCommandBlockData
	IDPlaceRuntimeBlockWithChestData
	IDPlaceRuntimeBlockWithChestDataAndUint32RuntimeID
	IDAssignDebugData
	IDPlaceBlockWithChestData
	IDPlaceBlockWithNBTData
	IDTerminate = 88
)

// Command ..
type Command interface {
	ID() uint16
	Name() string
	Marshal(io encoding.IO)
}

// CommandFunc ..
type CommandFunc func() Command

// BDumpCommandPool ..
func BDumpCommandPool() map[uint16]CommandFunc {
	return map[uint16]CommandFunc{
		1:  func() Command { return &CreateConstantString{} },
		5:  func() Command { return &PlaceBlockWithBlockStates{} },
		6:  func() Command { return &AddInt16ZValue0{} },
		7:  func() Command { return &PlaceBlock{} },
		8:  func() Command { return &AddZValue0{} },
		9:  func() Command { return &NoOperation{} },
		12: func() Command { return &AddInt32ZValue0{} },
		13: func() Command { return &PlaceBlockWithBlockStatesDeprecated{} },
		14: func() Command { return &AddXValue{} },
		15: func() Command { return &SubtractXValue{} },
		16: func() Command { return &AddYValue{} },
		17: func() Command { return &SubtractYValue{} },
		18: func() Command { return &AddZValue{} },
		19: func() Command { return &SubtractZValue{} },
		20: func() Command { return &AddInt16XValue{} },
		21: func() Command { return &AddInt32XValue{} },
		22: func() Command { return &AddInt16YValue{} },
		23: func() Command { return &AddInt32YValue{} },
		24: func() Command { return &AddInt16ZValue{} },
		25: func() Command { return &AddInt32ZValue{} },
		26: func() Command { return &SetCommandBlockData{} },
		27: func() Command { return &PlaceBlockWithCommandBlockData{} },
		28: func() Command { return &AddInt8XValue{} },
		29: func() Command { return &AddInt8YValue{} },
		30: func() Command { return &AddInt8ZValue{} },
		31: func() Command { return &UseRuntimeIDPool{} },
		32: func() Command { return &PlaceRuntimeBlock{} },
		33: func() Command { return &PlaceRuntimeBlockWithUint32RuntimeID{} },
		34: func() Command { return &PlaceRuntimeBlockWithCommandBlockData{} },
		35: func() Command { return &PlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID{} },
		36: func() Command { return &PlaceCommandBlockWithCommandBlockData{} },
		37: func() Command { return &PlaceRuntimeBlockWithChestData{} },
		38: func() Command { return &PlaceRuntimeBlockWithChestDataAndUint32RuntimeID{} },
		39: func() Command { return &AssignDebugData{} },
		40: func() Command { return &PlaceBlockWithChestData{} },
		41: func() Command { return &PlaceBlockWithNBTData{} },
		88: func() Command { return &Terminate{} },
	}
}
