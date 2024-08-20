package message
type OneBotFileType struct {
	Id string `json:"id,omitempty"`//文件 ID
	Name string `json:"name,omitempty"`//文件名
	Size int32 `json:"size,omitempty"`//文件大小
	Url string `json:"url,omitempty"`//下载链接
	Busid int64 `json:"busid,omitempty"`
}
type OneBotBot struct {
    Id int64 `json:"omitempty"`
    Avatar string `json:"omitempty"`
}
type OneBotReplyType struct {
	Action string  `json:"action,omitempty"`
	Params OneBotReplyParamsType `json:"params,omitempty"`
}
type OneBotReplyParamsType struct {
    Message_type string `json:"message_type,omitempty"`
    User_id int64 `json:"user_id,omitempty"`
    Group_id int64 `json:"group_id,omitempty"`
    Message string `json:"message,omitempty"`
}