package message
type MessageType struct {

	Retcode int32 `json:"retcode,omitempty"`//请求返回码
	Wording string `json:"wording,omitempty"`//请求返回信息
	Data struct {
	    Message_id int32 `json:"message_id,omitempty"`//消息 ID
	}
    Time    int64 `json:"time,omitempty"` //事件发生的时间戳
    Self_id int64 `json:"self_id,omitempty"`//收到事件的机器人 QQ 号
    Post_type string `json:"post_type,omitempty"`//上报类型
	Request_type string `json:"request_type,omitempty"`//请求类型
	Notice_type string `json:"notice_type,omitempty"`//通知类型
    Message_type string `json:"message_type,omitempty"` //消息类型
	Sub_type string `json:"sub_type,omitempty"`//消息子类型
    Message_id int32 `json:"message_id,omitempty"`//消息 ID
    User_id int64 `json:"user_id,omitempty"`//发送者 QQ 号
    Message Message `json:"message,omitempty"`//消息内容
    Raw_message string `json:"raw_message,omitempty"`//原始消息内容
    Font int32 `json:"font,omitempty"`//字体
	Group_id int64 `json:"group_id,omitempty"`//群号
	Operator_id int64 `json:"operator_id,omitempty"`//操作者 QQ 号,如果是主动退群，则和 user_id 相同
	Duration int64 `json:"duration,omitempty"`//禁言持续时间
	Target_id int64 `json:"target_id,omitempty"`//被戳者 QQ 号,运气王 QQ 号
	Honor_type string `json:"honor_type,omitempty"`//荣誉类型
	Comment string `json:"comment,omitempty"`//验证消息
	Flag string `json:"flag,omitempty"`//flag
	Anonymous struct {
	    Id string `json:"id,omitempty"`//匿名用户 ID
	    Name string `json:"name,omitempty"`//匿名用户名称
	    Flag string `json:"flag,omitempty"`//flag
	} `json:"anonymous,omitempty"`
    Sender struct {
        User_id int64 `json:"user_id,omitempty"`//发送者 QQ 号
        Nickname string `json:"nickname,omitempty"`//发送者昵称
        Card string `json:"card,omitempty"`//发送者名片
        Sex string `json:"sex,omitempty"`//发送者性别, male 或 female 或 unknown
        Age int32 `json:"age,omitempty"`//发送者年龄
		Role string `json:"role,omitempty"`//发送者角色, owner 或 admin 或 member
		Area string `json:"area,omitempty"`//地区
		Level string `json:"level,omitempty"`//等级
		Title string `json:"title,omitempty"`//头衔
    } `json:"sender,omitempty"`
	File struct {
		Id string `json:"id,omitempty"`//文件 ID
		Name string `json:"name,omitempty"`//文件名
		Size int32 `json:"size,omitempty"`//文件大小
		Url string `json:"url,omitempty"`//下载链接
		Busid int64 `json:"busid,omitempty"`
	} `json:"file,omitempty"`
}

// 定义 Message 类型
type Message []struct {
    Data struct {
        Text string `json:"text,omitempty"`
		Id string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`	
		Url string `json:"url,omitempty"`
		File string `json:"file,omitempty"`
		QQ string `json:"qq,omitempty"`
		Type string `json:"type,omitempty"`
		Title string `json:"title,omitempty"`
		Lat string `json:"lat,omitempty"`
		Lon string `json:"lon,omitempty"`
		Audio string `json:"audio,omitempty"`
		Data string `json:"data,omitempty"`

    } `json:"data,omitempty"`
    Type string `json:"type,omitempty"`
}
type EventType struct {
	 Type string `json:"type,omitempty"`
	Platform string `json:"platform,omitempty"`
	Post_type string `json:"post_type,omitempty"` // 请求类型
	Message_type string `json:"message_type,omitempty"` // 通知类型
	Request_type string `json:"request_type,omitempty"`
	Notice_type string `json:"notice_type,omitempty"`
	Time int64
	Sub_type string `json:"omitempty"`
	Message_id int32 `json:"omitempty"`
	User_id int64 `json:"omitempty"`
	Message string `json:"message,omitempty"`
	Msg string `json:"msg,omitempty"`
	Font int32 `json:"font,omitempty"`
	Raw_message Message `json:"raw_message,omitempty"`
	Group_id int64 `json:"group_id,omitempty"`
	User_avatar string `json:"user_avatar,omitempty"`
	Operator_id int64 `json:"operator_id,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	Target_id int64 `json:"target_id,omitempty"`
	Honor_type string `json:"honor_type,omitempty"`
	Comment string `json:"comment,omitempty"`
	Flag string `json:"flag,omitempty"`
	Anonymous struct {
	    Id string `json:"id,omitempty"`//匿名用户 ID
	    Name string `json:"name,omitempty"`//匿名用户名称
	    Flag string `json:"flag,omitempty"`//flag
	} `json:"anonymous,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Role string `json:"role,omitempty"`//发送者角色, owner 或 admin 或 member
	File struct {
		Id string `json:"id,omitempty"`//文件 ID
		Name string `json:"name,omitempty"`//文件名
		Size int32 `json:"size,omitempty"`//文件大小
		Url string `json:"url,omitempty"`//下载链接
		Busid int64 `json:"busid,omitempty"`
	}`json:"file,omitempty"`
	Bot struct {
    	Id int64 `json:"omitempty"`
    	Avatar string `json:"omitempty"`
	}`json:"bot,omitempty"`
}

//MESSAGES messages 消息事件
type NewEventType struct {
    Type string `json:"type,omitempty"`
	Platform string `json:"platform,omitempty"`
	Post_type string `json:"post_type,omitempty"` // 请求类型
	Message_type string `json:"message_type,omitempty"` // 通知类型
	Request_type string `json:"request_type,omitempty"`
	Notice_type string `json:"notice_type,omitempty"`
	Time int64
	Sub_type string `json:"omitempty"`
	Message_id int32 `json:"omitempty"`
	User_id int64 `json:"omitempty"`
	Message string `json:"message,omitempty"`
	Msg string `json:"msg,omitempty"`
	Font int32 `json:"font,omitempty"`
	Raw_message Message `json:"raw_message,omitempty"`
	Group_id int64 `json:"group_id,omitempty"`
	User_avatar string `json:"user_avatar,omitempty"`
	Operator_id int64 `json:"operator_id,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	Target_id int64 `json:"target_id,omitempty"`
	Honor_type string `json:"honor_type,omitempty"`
	Comment string `json:"comment,omitempty"`
	Flag string `json:"flag,omitempty"`
	Anonymous struct {
	    Id string `json:"id,omitempty"`//匿名用户 ID
	    Name string `json:"name,omitempty"`//匿名用户名称
	    Flag string `json:"flag,omitempty"`//flag
	} `json:"anonymous,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Role string `json:"role,omitempty"`//发送者角色, owner 或 admin 或 member
	File struct {
		Id string `json:"id,omitempty"`//文件 ID
		Name string `json:"name,omitempty"`//文件名
		Size int32 `json:"size,omitempty"`//文件大小
		Url string `json:"url,omitempty"`//下载链接
		Busid int64 `json:"busid,omitempty"`
	}`json:"file,omitempty"`
	Bot struct {
    	Id int64 `json:"omitempty"`
    	Avatar string `json:"omitempty"`
	}`json:"bot,omitempty"`
	Reply func(string) bool `json:"reply,omitempty"`
}