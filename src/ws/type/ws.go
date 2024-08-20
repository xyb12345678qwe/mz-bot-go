package WsType
import (
	Message "github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
)
type EventWsType struct {
	Type string `json:"type"`
	Data Message.EventType `json:"data"`
	Platform string `json:"platform"`
	Group_id int64 `json:"group_id"`
	Message_type string `json:"message_type"`
	Post_type string `json:"post_type"`
	Message string `json:"message"`
}
