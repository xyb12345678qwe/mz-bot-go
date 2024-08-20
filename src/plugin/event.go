package plugin
import (
  "fmt"
	"encoding/json"
Message "github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
MzWsType "github.com/xyb12345678qwe/mz-bot-go/src/ws/type"
"github.com/gorilla/websocket"
) 
func HandleEvent(event Message.EventType ,ws *websocket.Conn) Message.NewEventType {
	return Message.NewEventType{
        Type:          event.Type,
        Platform:      event.Platform,
        Post_type:     event.Post_type,
        Message_type:  event.Message_type,
        Request_type:  event.Request_type,
        Notice_type:   event.Notice_type,
        Time:          event.Time,
        Sub_type:      event.Sub_type,
        Message_id:    event.Message_id,
        User_id:       event.User_id,
        Message:       event.Message,
        Msg:           event.Msg,
        Font:          event.Font,
        Raw_message:   event.Raw_message,
        Group_id:      event.Group_id,
        User_avatar:   event.User_avatar,
        Operator_id:   event.Operator_id,
        Duration:      event.Duration,
        Target_id:     event.Target_id,
        Honor_type:    event.Honor_type,
        Comment:       event.Comment,
        Flag:          event.Flag,
        Anonymous:     event.Anonymous,
        Nickname:      event.Nickname,
        Role:          event.Role,
        File:          event.File,
        Bot:           event.Bot,
        Reply:  func (msg string) bool {
			    fmt.Println("发送消息",msg)  
			    e := MzWsType.EventWsType{
	        	Type: "reply",
				    Message_type: event.Message_type,
	        	Platform: event.Platform,
				    Group_id: event.Group_id,
				    Message: msg,
			    }
			    eventData, _ := json.Marshal(e)
	    	  ws.WriteMessage(websocket.TextMessage, eventData)
			    return true
		},
    }
}
