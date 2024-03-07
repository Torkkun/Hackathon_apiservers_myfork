package usecase

import (
	"app/database"
	"app/domain"
	"fmt"

	"github.com/google/uuid"
)

func CreateReplyUser(messageId, fromUserId string) error {
	replyID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	return database.CreateReplyUser(&database.ReplyUser{
		MessageID:  messageId,
		ReplyID:    replyID.String(),
		FromUserID: fromUserId,
	})
}

func GetReplyUser(messageId string) (*domain.ReplyUserList, error) {
	return database.GetReplyUser(messageId)
}

func CreateReplyMessage(replyID, fromUserId, message string) error {
	replyMessageID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	return database.CreateReplyMessage(&database.NewReplyMessage{
		ReplyID:        replyID,
		ReplyMessageID: replyMessageID.String(),
		ReplyText:      message,
		UserID:         fromUserId,
	})
}

func GetReplyMessage(replyId string) (*domain.ReplyMessageList, error) {
	return database.GetReplyMessage(replyId)
}

func CheckUserIdFromReplyUser(ruserID, replyID string) error {
	ownerID, userID, err := database.GetReplyUserByReplyMessageID(replyID)
	if err != nil {
		return err
	}
	if ruserID != ownerID && ruserID != userID {
		return fmt.Errorf("not exist UserID")
	}
	return nil
}
