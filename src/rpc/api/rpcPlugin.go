package rpc
import (
    "encoding/gob"
    "net/rpc"
    "bytes"
    "fmt"
    RpcType "github.com/xyb12345678qwe/mz-bot-go/src/rpc/type"
    Plugin "github.com/xyb12345678qwe/mz-bot-go/src/plugin"
)
type Client struct {}

func (c *Client) ReceiveMessage(message string, reply *string) error {
    // 处理服务端发送的消息
    fmt.Println("Received message from server:", message)
    *reply = "Message received"
    return nil
}

func RpcConnect() (*rpc.Client, error) {
    client, err := rpc.Dial("tcp", "127.0.0.1:9809")
    rpc.Register(&Client{})

    if err != nil {
        return nil, err
    }
    return client, nil
}

// func SendPlugin(client *rpc.Client) (RpcType.PluginResponse, error) {
//     // 编码
//     var network bytes.Buffer
//     enc := gob.NewEncoder(&network)
//     err := enc.Encode(Plugin.Plugins)
//     if err != nil {
//     fmt.Println("Error encoding Plugin.Plugins:", err)
//     return RpcType.PluginResponse{}, err
// }

//     // 提取 bytes.Buffer 中的数据到 []byte 中
//     pluginsData := network.Bytes()
//     stringMap := make(map[int]string)

//     for key, value := range Plugin.PluginsReg {
//         stringMap[key] = value.String()
//     }
//     fmt.Println(Plugin.Plugins)
//     fmt.Println(pluginsData)
//     req := RpcType.PluginRequest{
//         Plugins: pluginsData,
//         PluginsReg:stringMap,
//     }
//     var res RpcType.PluginResponse
//     err = client.Call("MzBotGo.AddPlugin", req, &res)
//     if err != nil {
//         return res, err
//     }
//     return res, nil
// }
