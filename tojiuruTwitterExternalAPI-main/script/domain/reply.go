package domain

type ReplyUserList []ReplyUser

type ReplyUser struct {
	ReplyID      string `json:"replyId"`
	MessageID    string `json:"messageId"`
	FromUserID   string `json:"fromUserId"`
	FromUserName string `json:"fromUserName"`
	CreatedAt    string `json:"createdAt"`
}

type PostReply struct {
	ReplyID   string `json:"replyId"`
	ReplyText string `json:"replyText"`
}

type Reply struct {
	ReplyID   string `json:"replyId"`
	FromUser  string `json:"fromUser"`
	CreatedAt string `json:"createdAt"`
}

type ReplyList []Reply

type ReplyMessage struct {
	ReplyID        string `json:"replyId"`
	ReplyMessageID string `json:"replyMessageId"`
	Message        string `json:"message"`
	UserID         string `json:"userId"`
	UserName       string `json:"username"`
	CreatedAt      string `json:"createdAt"`
}

type ReplyMessageList []*ReplyMessage
