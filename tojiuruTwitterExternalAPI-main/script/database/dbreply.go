package database

import (
	"app/domain"
	"database/sql"
)

type ReplyUser struct {
	MessageID  string
	ReplyID    string
	FromUserID string
}

// あれば無視無ければ作成
func CreateReplyUser(reply *ReplyUser) error {
	_, err := Db.Exec("INSERT OR IGNORE INTO replyuser(message_id, reply_id, from_user_id) values (?,?,?)",
		reply.MessageID, reply.ReplyID, reply.FromUserID)
	return err
}

func GetReplyUser(message_id string) (*domain.ReplyUserList, error) {
	rows, err := Db.Query(
		`SELECT
			message_id,
			reply_id,
			from_user_id,
			user.name,
			replyuser.created_at
		FROM
			replyuser
		INNER JOIN
			user 
		ON
			user.user_id = replyuser.from_user_id
		WHERE
			message_id = ?`,
		message_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var replyuser domain.ReplyUserList
	for rows.Next() {
		data, err := convertToReplyUser(rows)
		if err != nil {
			return nil, err
		}
		replyuser = append(replyuser, *data)
	}
	return &replyuser, nil
}

func convertToReplyUser(rows *sql.Rows) (*domain.ReplyUser, error) {
	replyu := domain.ReplyUser{}
	if err := rows.Scan(
		&replyu.ReplyID,
		&replyu.MessageID,
		&replyu.FromUserID,
		&replyu.FromUserName,
		&replyu.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &replyu, nil
}

type NewReplyMessage struct {
	ReplyID        string
	ReplyMessageID string
	ReplyText      string
	UserID         string
}

func CreateReplyMessage(replyMessage *NewReplyMessage) error {
	_, err := Db.Exec("INSERT INTO replymessage(reply_id, reply_message_id, reply_text, user_id) values (?,?,?,?)",
		replyMessage.ReplyID, replyMessage.ReplyMessageID, replyMessage.ReplyText, replyMessage.UserID)
	return err
}

func GetReplyMessage(replyID string) (*domain.ReplyMessageList, error) {
	rows, err := Db.Query(
		`SELECT
			reply_id,
			reply_message_id,
			reply_text,
			replymessage.user_id,
			user.name,
			replymessage.created_at
		FROM
			replymessage
		INNER JOIN
			user
		ON
			replymessage.user_id = user.user_id
		WHERE
			reply_id = ?`,
		replyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var replymessages domain.ReplyMessageList
	for rows.Next() {
		data, err := convertToReplyMessage(rows)
		if err != nil {
			return nil, err
		}
		replymessages = append(replymessages, data)
	}
	return &replymessages, nil
}

func convertToReplyMessage(rows *sql.Rows) (*domain.ReplyMessage, error) {
	replymessage := domain.ReplyMessage{}
	if err := rows.Scan(
		&replymessage.ReplyID,
		&replymessage.ReplyMessageID,
		&replymessage.Message,
		&replymessage.UserID,
		&replymessage.UserName,
		&replymessage.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &replymessage, nil
}

func GetReplyUserByReplyMessageID(messageID string) (string, string, error) {
	var ownerID, userID string
	row := Db.QueryRow(
		`SELECT
			replyuser.from_user_id,
			examination.user_id
		FROM
			replyuser
		INNER JOIN
			examination
		ON
			replyuser.message_id = examination.message_id
		WHERE
			reply_id = ?`, messageID)
	if err := row.Scan(&ownerID, &userID); err != nil {
		return "", "", err
	}
	return ownerID, userID, nil
}
