package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadStringFromPanel ..
func ReadStringFromPanel(langText string) (result string) {
	fmt.Print(langText)

	result, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	result = strings.ReplaceAll(result, "\n", "")
	result = strings.ReplaceAll(result, "\r", "")
	result = strings.ReplaceAll(result, "\t", "")

	return
}
