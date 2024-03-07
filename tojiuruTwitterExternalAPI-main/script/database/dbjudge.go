package database

import (
	"database/sql"
	"log"
)

type Judge struct {
	MessageId string
	UserId    string
	Judge     bool
	CreatedAt string
}

type JudgeCount struct {
	Sum   int
	Judge bool
}

// Judgeテーブル
func InsertJudge(judge *Judge) error {
	_, err := Db.Exec(
		`INSERT INTO 
			judge(message_id, user_id, judge, created_at)
		VALUES
			(?,?,?)`,
		judge.MessageId, judge.UserId, judge.Judge)
	return err
}

func CheckJudge(message_id string, uuid string) (*Judge, error) {
	row := Db.QueryRow(
		`SELECT
			*
		FROM 
			judge 
		WHERE 
			message_id = ? 
		AND 
			user_id = ?`, message_id, uuid)
	return convertToJudge(row)
}

// judgeテーブル内から条件に合うものをcountする 一審査テキストの良いね悪いね
func CountJudge(message_id string) (judgemap map[bool]int, err error) {
	judgemap = map[bool]int{true: 0, false: 0}
	query := "select count(*), judge FROM (SELECT * from judge where message_id = ?) GROUP BY judge"
	rows, err := Db.Query(query, message_id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var count int
		var judge bool
		if err := rows.Scan(&count, &judge); err != nil {
			log.Println(err)
			continue
		}
		judgemap[judge] = count
	}
	return
}

func convertToJudge(row *sql.Row) (*Judge, error) {
	judge := Judge{}
	if err := row.Scan(
		&judge.MessageId,
		&judge.UserId,
		&judge.Judge,
		&judge.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &judge, nil
}
