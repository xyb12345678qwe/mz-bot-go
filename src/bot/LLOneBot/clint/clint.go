package clint 
import (
	"github.com/gorilla/websocket"
	"fmt"
	Message "github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
	"encoding/json"
)
var Ws *websocket.Conn
func CreateClient(port string) *websocket.Conn {
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:" + port, nil)
    if err != nil {
		// 处理错误
		fmt.Println("[ws]连接到[LLOneBot]失败")
		return nil;
    }
    fmt.Println("[ws]连接到[LLOneBot]成功")
	Ws = conn
	return conn
}
func SendMsg(msg string,Message_type string ,Group_id int64) bool{
	data, _ := json.Marshal(Message.OneBotReplyType{
        Action: "send_msg",
        Params: Message.OneBotReplyParamsType{
            Message_type: Message_type,
            Group_id: Group_id,
            Message: msg,
        },
    })

    Ws.WriteMessage(websocket.TextMessage,data)
	return true
}