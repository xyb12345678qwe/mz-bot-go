package ws
import (
	"github.com/gorilla/websocket"
	"fmt"
)
var Ws *websocket.Conn
func ConnectWs() *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8719/ws", nil)
    if err != nil {
		fmt.Println("[plugin][ws]连接到[Mz-bot]失败")
		return nil;
    }
	Ws = conn
	return conn
}
 