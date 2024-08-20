package ws
import (
	"fmt"
	"time"
	"encoding/json"
   api "github.com/xyb12345678qwe/mz-bot-go/src/bot/LLOneBot/api"
   Message "github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
   Clint "github.com/xyb12345678qwe/mz-bot-go/src/bot/LLOneBot/clint"
)

func ConnectWs(port string) bool{
    clint := Clint.CreateClient(port)
    defer clint.Close()
	for {
        _, message, err := clint.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            // 可以根据错误类型进行处理，比如关闭连接、重新连接等
            break // 停止循环
        }
        if len(message) == 0 {
            continue
        }
        messageStr := string(message)
        var msg Message.MessageType
        err = json.Unmarshal([]byte(messageStr), &msg)
        if err != nil {
            fmt.Println("[LLOneBot][Error]:", err)
             continue // 继续循环
        }
        api.Conversion(msg)
        time.Sleep(100 * time.Millisecond) // 暂停0.1秒，控制消息获取频率
     //time.Millisecond 是1毫秒
    }
	return true;
}