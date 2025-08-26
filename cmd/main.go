package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/TriM-Organization/merry-memory/converter"
	"github.com/pterm/pterm"
)

//go:embed version
var versionInfo []byte

func init() {
	temp := strings.Split(string(versionInfo), "\n")[0]
	versionInfo = []byte(strings.TrimSpace(temp))
}

func main() {
	pterm.DefaultBox.Println(
		pterm.LightCyan("" +
			"                   " +
			"Merry Memory" + "\n" +
			"                     " +
			pterm.Sprintf("v%s", string(versionInfo)) + "\n" +
			"https://github.com/TriM-Organization/merry-memory",
		),
	)

	bdxPath := ReadStringFromPanel("请输入 BDX 文件路径: ")
	mcworldPath := ReadStringFromPanel("请输入 mcworld 文件路径: ")

	pterm.Info.Println("转换正在进行中, 请坐与放松...")
	err := converter.ConvertBDXToMCWorld(bdxPath, mcworldPath)
	if err != nil {
		pterm.Error.Printfln("处理时发生错误, 原因是: %v", err)
	}

	fmt.Println()
	ReadStringFromPanel(
		pterm.Success.Sprint("程序运行完成, 按 [回车] 以退出。"),
	)
}
