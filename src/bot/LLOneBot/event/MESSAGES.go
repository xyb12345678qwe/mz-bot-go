package event 


import (
	_"fmt"
	_"github.com/gorilla/websocket"
	_"encoding/json"
    Message "github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
	_"github.com/xyb12345678qwe/mz-bot-go/src/bot/LLOneBot/clint"
	MzWs "github.com/xyb12345678qwe/mz-bot-go/src/ws/api"
	MzWsType "github.com/xyb12345678qwe/mz-bot-go/src/ws/type"
)

func MESSAGES(message Message.MessageType) Message.EventType {
	e :=Message.EventType{
		Type: "MESSAGES",
		Platform: "LLOneBot",
		Post_type: message.Post_type,
		Message_type: message.Message_type,
		Request_type: "",
		Notice_type: "",
		Time: message.Time ,
		Sub_type: message.Sub_type,
		Message_id: message.Message_id,
		User_id: message.User_id,
		User_avatar: `http://q.qlogo.cn/headimg_dl?dst_uin=`+ string(message.User_id) + `&spec=640&img_type=jpg`,
		Message: message.Raw_message,
		Msg: message.Raw_message,
		Raw_message: message.Message,
		Font: message.Font,
		Group_id: message.Group_id,
		Operator_id: 0,
		Duration: 0,
		Target_id: 0,
		Honor_type: "",
		Comment: "",
		Flag: "",
		Anonymous: message.Anonymous,
		Nickname: message.Sender.Nickname,
		Role: message.Sender.Role,
		File:Message.OneBotFileType{
			Id:"",
			Name: "",
			Size: 0,
			Url: "",
			Busid: 0,
		},
		Bot:Message.OneBotBot{
			Id:message.Self_id,
			Avatar: `http://q.qlogo.cn/headimg_dl?dst_uin=`+ string(message.Self_id) + `&spec=640&img_type=jpg`,
		},
		// Reply: func(msg string) bool{
		// 	data, _ := json.Marshal(Message.OneBotReplyType{
        // 		Action: "send_msg",
        // 		Params: Message.OneBotReplyParamsType{
        //     		Message_type: message.Post_type,
        //     		Group_id: message.Group_id,
        //     		Message: msg,
        // 		},
    	// 	})
    	// 	clint.Ws.WriteMessage(websocket.TextMessage,data)
		// 	return true
		// },
	}
	MzWs.SendEvent(MzWsType.EventWsType{
		Type: "MESSAGES",
		Data: e,
	})
	return e
}