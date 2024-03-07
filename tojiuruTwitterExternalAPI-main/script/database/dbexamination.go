package database

import (
	"database/sql"
	"log"
)

type Examination struct {
	MessageId string
	Message   string
	People    int
	CreatedAt string
	Deadline  string
	UserId    string
	UserName  string
}

type ExaminationList []*Examination

type TweetS struct {
	Tweet_ID    string
	Tweet_text  string
	AccessToken string
	SecretToken string
}

type TweetList []TweetS

func CreateExamination(examination *Examination) error {
	_, err := Db.Exec(
		`INSERT INTO
			examination(message_id, message, people_num, user_id, deadline)
		VALUES
			(?,?,?,?,?)`,
		examination.MessageId, examination.Message, examination.People, examination.UserId, examination.Deadline)
	return err
}

// Tx
func CreateExaminationTx(tx *sql.Tx, examination *Examination) error {
	_, err := tx.Exec(
		`INSERT INTO
			examination(message_id, message,  people_num, user_id ,deadline )
		VALUES
			(?,?,?,?,?)`,
		examination.MessageId, examination.Message, examination.People, examination.UserId, examination.Deadline)
	return err
}

func SelectExamination() (*ExaminationList, error) {
	rows, err := Db.Query(
		`SELECT
			message_id,
			message,
			people_num,
			examination.user_id,
			user.name,
			deadline,
			examination.created_at
		FROM
			examination
		INNER JOIN
			user
		ON
			user.user_id = examination.user_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var examinations ExaminationList
	for rows.Next() {
		examination, err := convertToExamination(rows)
		if err != nil {
			return nil, err
		}
		examinations = append(examinations, examination)
	}
	return &examinations, nil
}

func FindUserIDFromExaminationByMessageID(id string) string {
	var uidfromdb string
	Db.QueryRow(
		`SELECT
			examination.user_id
		FROM
			examination
		WHERE 
			message_id = ?`, id).Scan(&uidfromdb)
	return uidfromdb
}

// 現在時刻より前のデータを抜き出す
// *修正deadlineがどうのこうのはDBでやらない
func Checkdeadline() (TweetList, error) {
	var tweetlist []TweetS
	rows, err := Db.Query(
		`SELECT
			tweet_id, tweet_text, accesstoken, secrettoken
		FROM 
			examination
		WHERE 
			deadline <= datetime(CURRENT_TIMESTAMP, 'localtime') 
		AND deadline != ''`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tweets TweetS
		if err := rows.Scan(&tweets.Tweet_ID, &tweets.Tweet_text, &tweets.AccessToken, &tweets.SecretToken); err != nil {
			log.Println(err)
			return nil, err
		}
		tweetlist = append(tweetlist, tweets)
	}
	return tweetlist, nil
}

func DeleteExamination(tweet_id string) (err error) {
	statement, err := Db.Prepare("DELETE FROM examination where message_id = ?")
	if err != nil {
		return
	}
	defer statement.Close()
	_, err = statement.Exec(tweet_id)
	return
}

func convertToExamination(rows *sql.Rows) (*Examination, error) {
	exam := Examination{}
	if err := rows.Scan(
		&exam.MessageId,
		&exam.Message,
		&exam.People,
		&exam.UserId,
		&exam.UserName,
		&exam.CreatedAt,
		&exam.Deadline); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &exam, nil
}
