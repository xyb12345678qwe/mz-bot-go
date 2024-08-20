package api
import (
"fmt"
LLtype "github.com/xyb12345678qwe/mz-bot-go/src/bot/LLOneBot/type"
Message "github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
Event "github.com/xyb12345678qwe/mz-bot-go/src/bot/LLOneBot/event"
"reflect"
)

var v = reflect.ValueOf(LLtype.NoticeType)
var v2 = reflect.ValueOf(LLtype.RequestType)
func Conversion(message Message.MessageType){
	if message.Post_type == "message" {
		Event.MESSAGES(message)
		sender := message.Sender
		fmt.Printf("[LLOneBot][Message][%s(%d)]:%s\n",sender.Nickname,sender.User_id,message.Raw_message)
	} else if message.Post_type == "notice" {
		notice_type := message.Notice_type
		fmt.Printf("[LLOneBot][Notice][%s]:%s\n",notice_type,v.FieldByName(notice_type))
	} else if message.Post_type == "meta_event" {
		fmt.Printf("[LLOneBot][meta_event]:心跳检测\n")
	} else if message.Post_type == "request" {
		request_type := message.Request_type
		fmt.Printf("[LLOneBot][request][%s]:%s\n",request_type,v2.FieldByName(request_type))
	}
}
//  {"self_id":3135864774,
//  "user_id":4114247,
//  "time":1723739872,
//  "message_id":-2147479330,
//  "real_id":-2147479330,
//  "message_seq":-2147479330,"
//  message_type":"group",
//  "sender":{"user_id":4114247,"nickname":"沉铂纹","card":"药尊者","role":"admin"},
//  "raw_message":"😂",
//  "font":14,
//  "sub_type":"normal",
//  "message":[{"
// data":{"text":"😂"},"type":"text"}],
// "message_format":"array",
// "post_type":"message",
// "group_id":476189165
// }
