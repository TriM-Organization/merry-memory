package utils

import (
	"bytes"
	"fmt"
	"os"

	"github.com/TriM-Organization/merry-memory/command"
	"github.com/TriM-Organization/merry-memory/protocol/encoding"
	"github.com/andybalholm/brotli"
)

// ReadBDXFileInfo ..
func ReadBDXFileInfo(path string) (authorName string, reader *encoding.Reader, fileCloser func() error, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("ReadBDXFileInfo: %v", r)
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		return "", nil, nil, fmt.Errorf("ReadBDXFileInfo: %v", err)
	}

	outsideHeader := make([]byte, 3)
	readBytes, _ := file.Read(outsideHeader)
	if readBytes < 3 {
		return "", nil, nil, fmt.Errorf("ReadBDXFileInfo: Not a valid BDX file (Outside header too short)")
	}
	if string(outsideHeader) != "BD@" {
		return "", nil, nil, fmt.Errorf(`ReadBDXFileInfo: File outside header %#v not match magic string "BD@"`, string(outsideHeader))
	}

	brotliReader := brotli.NewReader(file)
	insideHeader := make([]byte, 4)
	readBytes, _ = brotliReader.Read(insideHeader)
	if readBytes < 4 {
		return "", nil, nil, fmt.Errorf("ReadBDXFileInfo: Not a valid BDX file (Inside header too short)")
	}
	if !bytes.Equal(insideHeader, append([]byte("BDX"), 0)) {
		return "", nil, nil, fmt.Errorf(`ReadBDXFileInfo: File inside header %#v not match magic string "BDX\x00"`, insideHeader)
	}

	reader = encoding.NewReader(NewGeneralReader(brotliReader))
	reader.CString(&authorName)
	return authorName, reader, file.Close, nil
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
