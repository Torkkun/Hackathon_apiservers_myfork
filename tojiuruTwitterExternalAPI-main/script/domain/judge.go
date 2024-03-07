package domain

type PostJudge struct {
	Judge     bool   `json:"judge"`
	MessageId string `json:"message_id"`
}
