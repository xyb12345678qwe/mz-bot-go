package main

import (
	"fmt"
	"encoding/json"
	_"mz-bot-go/plugins/test/src/apps"
	Plugin "github.com/xyb12345678qwe/mz-bot-go/src/plugin"
	MzWs"github.com/xyb12345678qwe/mz-bot-go/src/ws/api"
	MzWsType"github.com/xyb12345678qwe/mz-bot-go/src/ws/type"
)

func main() {
	fmt.Println("test plugin load")
	clint:= MzWs.ConnectWs()

	defer clint.Close()
	for {
		 _, message, err := clint.ReadMessage()
		 if err != nil {
            fmt.Println("[Error][read][message]:", err)
            // 可以根据错误类型进行处理，比如关闭连接、重新连接等
            break // 停止循环
        }
		if len(message) == 0 {
            continue
        }
		messageStr := string(message)
		var msg MzWsType.EventWsType
        err = json.Unmarshal([]byte(messageStr), &msg)
        if err != nil {
            fmt.Println("[plugin][Error]:", err)
            continue // 继续循环
        }
		newMsg := Plugin.HandleEvent(msg.Data,clint)
		// fmt.Println("[plugin][newMsg]:",newMsg)
		Plugin.Response(newMsg)
	}
}
