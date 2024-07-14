package main

import (
	"fmt"
)

// 插件接口
type Plugin interface {
	Process(data string) string
}

// 具体插件实现
type UppercasePlugin struct{}

func (u *UppercasePlugin) Process(data string) string {
	return fmt.Sprintf("UppercasePlugin: %s", data)
}

type ReversePlugin struct{}

func (r *ReversePlugin) Process(data string) string {
	return fmt.Sprintf("ReversePlugin: %s", reverseString(data))
}

// 反转字符串函数
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 插件组合结构
type PluginSystem struct {
	plugins []Plugin
}

// 添加插件到系统
func (ps *PluginSystem) AddPlugin(plugin Plugin) {
	ps.plugins = append(ps.plugins, plugin)
}

// 使用插件系统处理数据
func (ps *PluginSystem) ProcessData(data string) {
	for _, plugin := range ps.plugins {
		result := plugin.Process(data)
		fmt.Println(result)
		// 在实际应用中，可以将结果传递给下一个插件或进行其他操作
	}
}

func main() {
	// 创建插件实例
	uppercasePlugin := &UppercasePlugin{}
	reversePlugin := &ReversePlugin{}

	// 创建插件系统并添加插件
	pluginSystem := PluginSystem{}
	pluginSystem.AddPlugin(uppercasePlugin)
	pluginSystem.AddPlugin(reversePlugin)

	// 定义输入数据
	inputData := "Hello, World!"

	// 使用插件系统处理数据
	pluginSystem.ProcessData(inputData)
	pluginSystem.ProcessData("Goodbye, World!")
}
