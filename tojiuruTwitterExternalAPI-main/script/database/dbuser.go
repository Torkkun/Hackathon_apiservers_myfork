package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	UserId      string
	Name        string
	Password    string
	AccessToken string
	SecretToken string
	CreatedAt   string //sqliteではstringとして扱う為（後に不都合があれば変換させる）
}

type UserList []*User

//ユーザーデータを作成
func CreateUser(user *User) (err error) {
	statement, err := Db.Prepare("INSERT INTO user(user_id, name, password) values (?,?,?)")
	if err != nil {
		return
	}
	defer statement.Close()
	uid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	_, err = statement.Exec(uid, user.Name, Encrypt(user.Password))
	return
}

func GetUsers() (*UserList, error) {
	rows, err := Db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users UserList
	for rows.Next() {
		user, err := convertToUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, err

}

//名前から情報を取得
func UserByName(name string) (*User, error) {
	var userID string
	var pass string
	if err := Db.QueryRow("SELECT user_id, password FROM user WHERE name = ?", name).Scan(&userID, &pass); err != nil {
		return nil, err
	}
	return &User{
		UserId:   userID,
		Password: pass,
	}, nil
}

// examination登録の際にアップデートをする処理を入れる（未完了）
func UpdateTokenByUserId(user *User) error {
	_, err := Db.Exec(
		"UPDATE user SET accesstoken = ?, secrettoken = ? WHERE user_id = ?",
		user.AccessToken, user.SecretToken, user.UserId)
	return err
}

func convertToUser(rows *sql.Rows) (*User, error) {
	user := User{}
	if err := rows.Scan(
		&user.UserId,
		&user.Name,
		&user.Password,
		&user.AccessToken,
		&user.SecretToken,
		&user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

