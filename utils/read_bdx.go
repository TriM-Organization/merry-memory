package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/TriM-Organization/merry-memory/command"
	"github.com/TriM-Organization/merry-memory/protocol/encoding"
	"github.com/andybalholm/brotli"
)

// ReadBDXFileInfo ..
func ReadBDXFileInfo(path string) (authorName string, reader *encoding.Reader, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("ReadBDXFileInfo: %v", r)
		}
	}()

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", nil, fmt.Errorf("ReadBDXFileInfo: %v", err)
	}

	if len(fileBytes) < 3 {
		return "", nil, fmt.Errorf("ReadBDXFileInfo: Not a valid BDX file (Outside header too short)")
	}
	if string(fileBytes[0:3]) != "BD@" {
		return "", nil, fmt.Errorf(`ReadBDXFileInfo: File outside header %#v not match magic string "BD@"`, string(fileBytes[0:3]))
	}

	buf := bytes.NewBuffer(fileBytes[3:])
	brotliReader := brotli.NewReader(buf)
	fileBytes, err = io.ReadAll(brotliReader)
	if err != nil {
		return "", nil, fmt.Errorf("ReadBDXFileInfo: %v", err)
	}

	if len(fileBytes) < 4 {
		return "", nil, fmt.Errorf("ReadBDXFileInfo: Not a valid BDX file (Inside header too short)")
	}
	if !bytes.Equal(fileBytes[0:4], append([]byte("BDX"), 0)) {
		return "", nil, fmt.Errorf(`ReadBDXFileInfo: File inside header %#v not match magic string "BDX\x00"`, fileBytes[0:4])
	}
	fileBytes = fileBytes[4:]

	buf = bytes.NewBuffer(fileBytes)
	reader = encoding.NewReader(buf)
	reader.CString(&authorName)

	return authorName, reader, nil
}

// ReadBDXCommand ..
func ReadBDXCommand(reader *encoding.Reader) (cmd command.Command, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("ReadBDXCommand: %v", r)
		}
	}()

	var commandID uint8
	reader.Uint8(&commandID)

	commandFunc, ok := command.BDumpCommandPool[uint16(commandID)]
	if !ok {
		return nil, fmt.Errorf("ReadBDXCommand: Command ID %d is not supported", commandID)
	}

	cmd = commandFunc()
	cmd.Marshal(reader)

	return cmd, nil
}
