package utils

import (
	"strconv"
	"strings"
)

// "color":"orange" [legacy]
// or
// "color"="orange" [current]
const (
	BlockStatesDefaultSeparator = "="
	BlockStatesLegacySeparator  = ":"
)

// ParseBlockStatesString ..
func ParseBlockStatesString(blockStatesString string, isLegacy bool) (result map[string]any) {
	if len(blockStatesString) < 2 {
		return make(map[string]any)
	}
	if blockStatesString[0] != '[' || blockStatesString[len(blockStatesString)-1] != ']' {
		return make(map[string]any)
	}

	separator := BlockStatesDefaultSeparator
	if isLegacy {
		separator = BlockStatesLegacySeparator
	}
	blockStatesString = blockStatesString[1 : len(blockStatesString)-1]
	result = make(map[string]any)

	for state := range strings.SplitSeq(blockStatesString, ",") {
		state = strings.TrimSpace(state)

		temp := strings.Split(state, separator)
		if len(temp) != 2 {
			continue
		}

		key := strings.ReplaceAll(temp[0], `"`, "")
		if len(temp[1]) < 1 {
			continue
		}
		value := temp[1]

		switch value[0] {
		case '"':
			result[key] = strings.ReplaceAll(value, `"`, "")
		case 't', 'T':
			result[key] = byte(1)
		case 'f', 'F':
			result[key] = byte(0)
		default:
			val, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				continue
			}
			result[key] = int32(val)
		}
	}

	return
}
