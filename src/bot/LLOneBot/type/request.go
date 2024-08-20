package LLtype
type requestType struct {
    friend string
	group string
}
var RequestType = requestType{
    friend: "加好友请求",
	group: "加群请求/邀请",
}