package main

import (
	"fmt"
	"os"
	_"strconv"
	LLOneBot "github.com/xyb12345678qwe/mz-bot-go/src/bot/LLOneBot/ws"
	_ "github.com/xyb12345678qwe/mz-bot-go/src/file"
	Plugin "github.com/xyb12345678qwe/mz-bot-go/src/plugin"
	utils "github.com/xyb12345678qwe/mz-bot-go/src/utils"
	MzWs "github.com/xyb12345678qwe/mz-bot-go/src/ws/api"
)

func main() {
	fmt.Println("Mz-bot is running")
	go func() {
		MzWs.CreateWs()
	}()
	argsWithoutProg := os.Args[1:] // 去除第一个元素，即程序的路径
	Plugin.PluginLoad("plugins")
	
	//后台运行
	// maxKey := Plugin.FindMaxKey(Plugin.Plugins)

	// 将整数转换为字符串，然后再进行拼接操作
	// result := "[加载插件]" + strconv.Itoa(maxKey)

	// fmt.Println(result)
	//获取要启动机器人的数组
	if utils.Includes(utils.ToInterfaceSlice(argsWithoutProg), "LLOneBot") {
		fmt.Println(`[LLOneBot] start`)
		LLOneBot.ConnectWs("3001")
	}
	
	// config , _ := file.ReadFileByYAML("botConfig");
	// value,_ := config["test"].(map[string]interface{})
	// value["ntqq"]= map[string]interface{}{}
	// file.WriteFileByYAML(config, "botConfig")
	// fmt.Println(value)
}
