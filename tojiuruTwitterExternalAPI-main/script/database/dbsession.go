package database

type Session struct {
	UserId    string
	CreatedAt string
}

//新しいセッションを作成
func CreateSession(userID string) (err error) {
	statement, err := Db.Prepare(
		`INSERT INTO 
			sessions(user_id) VALUES (?) on conflict(user_id) do update SET created_at = datetime(CURRENT_TIMESTAMP, 'localtime')`)
	if err != nil {
		return
	}
	defer statement.Close()
	if err != nil {
		return
	}
	_, err = statement.Exec(userID)
	return
}

func DeleteByUserID(session *Session) error {
	statement, err := Db.Prepare("DELETE FROM sessions where user_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(session.UserId)
	return err
}
