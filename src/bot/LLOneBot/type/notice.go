package LLtype
type noticeType struct {
    group_upload string
	group_admin string
	group_decrease string
	group_increase string
	group_ban string
	friend_add string
	group_recall string
	friend_recall string
	notify string
}
var NoticeType = noticeType{
    group_upload: "群文件上传",
	group_admin: "群管理员变动",
	group_decrease: "群成员减少",
	group_increase: "群成员增加",
	group_ban: "群禁言",
	friend_add: "好友添加",
	group_recall: "群消息撤回",
	friend_recall: "好友消息撤回",
	notify: "群内戳一戳,群红包运气王",
}