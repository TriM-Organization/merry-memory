package command

import "github.com/TriM-Organization/merry-memory/protocol/encoding"

const (
	IDCreateConstantString uint16 = iota + 1
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
var BDumpCommandPool = map[uint16]CommandFunc{
	IDCreateConstantString:                  func() Command { return &CreateConstantString{} },
	IDPlaceBlockWithBlockStates:             func() Command { return &PlaceBlockWithBlockStates{} },
	IDAddInt16ZValue0:                       func() Command { return &AddInt16ZValue0{} },
	IDPlaceBlock:                            func() Command { return &PlaceBlock{} },
	IDAddZValue0:                            func() Command { return &AddZValue0{} },
	IDNOP:                                   func() Command { return &NoOperation{} },
	IDAddInt32ZValue0:                       func() Command { return &AddInt32ZValue0{} },
	IDPlaceBlockWithBlockStatesDeprecated:   func() Command { return &PlaceBlockWithBlockStatesDeprecated{} },
	IDAddXValue:                             func() Command { return &AddXValue{} },
	IDSubtractXValue:                        func() Command { return &SubtractXValue{} },
	IDAddYValue:                             func() Command { return &AddYValue{} },
	IDSubtractYValue:                        func() Command { return &SubtractYValue{} },
	IDAddZValue:                             func() Command { return &AddZValue{} },
	IDSubtractZValue:                        func() Command { return &SubtractZValue{} },
	IDAddInt16XValue:                        func() Command { return &AddInt16XValue{} },
	IDAddInt32XValue:                        func() Command { return &AddInt32XValue{} },
	IDAddInt16YValue:                        func() Command { return &AddInt16YValue{} },
	IDAddInt32YValue:                        func() Command { return &AddInt32YValue{} },
	IDAddInt16ZValue:                        func() Command { return &AddInt16ZValue{} },
	IDAddInt32ZValue:                        func() Command { return &AddInt32ZValue{} },
	IDSetCommandBlockData:                   func() Command { return &SetCommandBlockData{} },
	IDPlaceBlockWithCommandBlockData:        func() Command { return &PlaceBlockWithCommandBlockData{} },
	IDAddInt8XValue:                         func() Command { return &AddInt8XValue{} },
	IDAddInt8YValue:                         func() Command { return &AddInt8YValue{} },
	IDAddInt8ZValue:                         func() Command { return &AddInt8ZValue{} },
	IDUseRuntimeIDPool:                      func() Command { return &UseRuntimeIDPool{} },
	IDPlaceRuntimeBlock:                     func() Command { return &PlaceRuntimeBlock{} },
	IDPlaceRuntimeBlockWithUint32RuntimeID:  func() Command { return &PlaceRuntimeBlockWithUint32RuntimeID{} },
	IDPlaceRuntimeBlockWithCommandBlockData: func() Command { return &PlaceRuntimeBlockWithCommandBlockData{} },
	IDPlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID: func() Command { return &PlaceRuntimeBlockWithCommandBlockDataAndUint32RuntimeID{} },
	IDPlaceCommandBlockWithCommandBlockData:                   func() Command { return &PlaceCommandBlockWithCommandBlockData{} },
	IDPlaceRuntimeBlockWithChestData:                          func() Command { return &PlaceRuntimeBlockWithChestData{} },
	IDPlaceRuntimeBlockWithChestDataAndUint32RuntimeID:        func() Command { return &PlaceRuntimeBlockWithChestDataAndUint32RuntimeID{} },
	IDAssignDebugData:         func() Command { return &AssignDebugData{} },
	IDPlaceBlockWithChestData: func() Command { return &PlaceBlockWithChestData{} },
	IDPlaceBlockWithNBTData:   func() Command { return &PlaceBlockWithNBTData{} },
	IDTerminate:               func() Command { return &Terminate{} },
}
