package plugin
import (
	"fmt"
	Message "github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
	"regexp"
	"strings"
)

var Plugins = make(map[int]func(Message.NewEventType)) //存基础的Messages消息
var PluginsReg = make(map[int]*regexp.Regexp) //正则
//EventType 事件类型
// reg 正则表达式
// plugin 处理函数
func RegisterPlugin(EventType string ,reg *regexp.Regexp, plugin func(Message.NewEventType)) {
	fmt.Println("[RegisterPlugin][注册指令]")
	switch EventType {
		case "MESSAGES":
			key := FindMaxKey(Plugins)
			Plugins[key + 1] = plugin
			PluginsReg[key + 1] = reg
		//不做任何处理
		default:
			panic("Not support event type")
	}
}

// findMaxKey 遍历 map 并返回最大的键
func FindMaxKey(m map[int]func(Message.NewEventType)) int {
    var max int
    first := true
    for key := range m {
        if first || key > max {
            max = key
            first = false
        }
    }
	if max == 0 && !first { // Correction for the logical condition check
		max = 0
	}
    return max
}

func Response(e Message.NewEventType){
	// fmt.Println("[Response][收到消息]")
	if e.Type == "MESSAGES" {
		for key, value := range PluginsReg {
			if value.MatchString(strings.TrimSpace(e.Message)) {
				fmt.Println("[Response][匹配到指令]")
				Plugins[key](e)
			}
		}
	}
}
