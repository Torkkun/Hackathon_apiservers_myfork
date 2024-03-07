package database

import (
	"log"
	"time"

	"github.com/google/uuid"
	"yandereca.tech/yandereca/domain"
)

type UserDataRepository struct {
	SqlHandler
}

func (repo *UserDataRepository) Store(userdata domain.UserData) (id string, err error) {
	if len(userdata.Id) == 0 {
		uid, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		uid_str := uid.String()
		userdata.Id = uid_str
	}

	_, err = repo.Execute(
		"INSERT INTO userdata(id, name, email, token, refresh_token, google_uid) values($1,$2,$3,$4,$5,$6) ON CONFLICT ON CONSTRAINT userdata_pkey DO UPDATE SET name=$2, email=$3, token=$4, refresh_token=$5, google_uid=$6",
		userdata.Id, userdata.Name, userdata.Email, userdata.Token, userdata.RefreshToken, userdata.GoogleUid,
	)
	if err != nil {
		return "", err
	}
	id = userdata.Id
	return
}

func (repo *UserDataRepository) FindById(identifier string) (userdataList domain.UserDataList, err error) {
	rows, err := repo.Query("SELECT * FROM userdata WHERE id = $1", identifier)

	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var email string
		var token string
		var refresh_token string
		var google_uid string
		var created_at time.Time
		var updated_at time.Time
		if err := rows.Scan(&id, &name, &email, &token, &refresh_token, &google_uid, &created_at, &updated_at); err != nil {
			log.Println(err)
			continue
		}
		userdata := domain.UserData{
			Id:           id,
			Name:         name,
			Email:        email,
			Token:        token,
			RefreshToken: refresh_token,
			GoogleUid:    google_uid,
		}

		userdataList = append(userdataList, userdata)
	}
	return
}

func (repo *UserDataRepository) Remove(identifier string) (err error) {
	_, err = repo.Execute("DELETE FROM userdata WHERE id = $1", identifier)
	if err != nil {
		return err
	}
	return
}
