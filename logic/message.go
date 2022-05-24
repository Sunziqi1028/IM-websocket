package logic

type Message struct {
	From    uint64
	To      []uint64
	Content string
	MsgType string //  消息类型 room ：聊天室 | radio：广播  | orient：定向
}
