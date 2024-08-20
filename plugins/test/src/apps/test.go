package apps

import (
	Plugin "github.com/xyb12345678qwe/mz-bot-go/src/plugin"
	Message "github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
	"regexp"
)
func init() {
	regexPattern := `^(#|/)?test$`
	regex, _ := regexp.Compile(regexPattern)
	Plugin.RegisterPlugin("MESSAGES", regex, func(e Message.NewEventType) {
		e.Reply("Hello World")
	})
}