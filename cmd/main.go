package main

import (
	"fmt"

	"github.com/TriM-Organization/merry-memory/converter"
)

func main() {
	bdxPath := ReadStringFromPanel("请输入 BDX 文件路径: ")
	mcworldPath := ReadStringFromPanel("请输入 mcworld 文件路径: ")

	fmt.Println("转换正在进行中, 请坐与放松...")
	err := converter.ConvertBDXToMCWorld(bdxPath, mcworldPath)
	if err != nil {
		fmt.Printf("处理时发生错误, 原因是: %v", err)
	}

	fmt.Println()
	ReadStringFromPanel("程序运行完成, 按 [回车] 以退出。")
}
