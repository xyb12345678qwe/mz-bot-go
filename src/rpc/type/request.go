package RpcType
import ( 
	_"bytes"
	_"regexp"
	_"github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
 )
// 定义 RPC 调用的请求结构体
type PluginRequest struct {
    Plugins []byte
	PluginsReg map[int]string
}
