package rpc
import (
    "encoding/gob"
    "net"
    "net/rpc"
    "fmt"
    "bytes"
    "regexp"
    Message"github.com/xyb12345678qwe/mz-bot-go/src/bot/message"
    RpcType "github.com/xyb12345678qwe/mz-bot-go/src/rpc/type"
    Plugin "github.com/xyb12345678qwe/mz-bot-go/src/plugin"
)
type ClientInfo struct {
    Address string
}
var clients = make(map[string]ClientInfo)

// 注册客户端的方法
func (s *MzBotGo) RegisterClientAddress(address string) {
    clients[address] = ClientInfo{address}
}

// 创建一个结构体用于实现 RPC 方法
type MzBotGo struct{}
func sendToClient(address string, message string) error {
    client, err := rpc.Dial("tcp", address)
    if err != nil {
        return err
    }
    var reply string
    err = client.Call("Client.ReceiveMessage", message, &reply)
    if err != nil {
        return err
    }
    return nil
}

func CreateRpc() {
    // 实例化服务对象
    MzBotGo := new(MzBotGo)

    // 将服务对象注册为 RPC 服务
    rpc.Register(MzBotGo)

    // 创建 TCP 监听
    listener, err := net.Listen("tcp", ":9809")
    if err != nil {
        fmt.Println("[rpc][Listen][error]:", err)
    }

    defer listener.Close()

    fmt.Println("[rpc][server][port]:9809")

    // 循环等待客户端连接
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("[rpc][accept][error]:", err)
        }
        
        go rpc.ServeConn(conn) // 处理客户端请求
    }
}