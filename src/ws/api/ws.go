package ws
import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	MzWsType "github.com/xyb12345678qwe/mz-bot-go/src/ws/type"
	LLOneBotClint "github.com/xyb12345678qwe/mz-bot-go/src/bot/LLOneBot/clint"
)
var clients = make(map[*websocket.Conn]bool) // 客户端连接映射
var broadcast = make(chan []byte) // 广播消息通道

var upgrader = websocket.Upgrader{}
func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	clients[conn] = true

	for {
		_, message , err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error: %v", err)
			delete(clients, conn)
			break
		}
		if len(message) == 0 {
            continue
        }
		// messageStr := string(message)
		var msg MzWsType.EventWsType
        // err = json.Unmarshal([]byte(messageStr), &msg)
        err = json.Unmarshal(message, &msg)
        if err != nil {
            fmt.Println("[plugin][Error]:", err)
            continue // 继续循环
        }
		fmt.Println("[ws][server][receive]:", msg.Message)
		if msg.Type == "reply"{
			if msg.Platform == "LLOneBot" { 
				LLOneBotClint.SendMsg(msg.Message , msg.Message_type, msg.Group_id)
			}
		}
		// broadcast <- msg 不进行发送管道
	}
}
//从管道中获取消息，并广播给所有客户端
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("Error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
//发送event事件到所有客户端
func SendEvent(data MzWsType.EventWsType) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data: %v", err)
		return
	}
	broadcast <- jsonData
}

func CreateWs() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	fmt.Println("[ws][server][port]:8719")
	err := http.ListenAndServe(":8719", nil)
	if err != nil {
		fmt.Println(err)
	}
}