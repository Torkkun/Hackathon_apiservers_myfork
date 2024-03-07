package todo

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"yandereca.tech/yandereca/domain"
)

type userJson struct {
	UserId   string `json:"uid"`
	FullName string `json:"fullname"`
	SurName  string `json:"surname"`
	Name     string `json:"name"`
}

func newJwtCreate(userInfo *userJson) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userInfo})
	// 思い返してみると何も暗号化の秘密鍵設定してなかった
	// hamcはあれなのでEdDSAで暗号化した方が良いかも
	// 秘密鍵はどこか良い保管場所が欲しい
	var hmacSampleSecret []byte
	signedToken, err := jwtToken.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

type progress struct {
	toDay  float64
	toWeek float64
}

//一桁の場合0パディングする用
func formatTime(due time.Time) string {
	// FormatはAPIによって違いそうなだが面倒なので後々考える
	time := due.Format("2006-01-02T00:00:00.000Z")
	return time
}

// 今日の進捗を計算
func toDayProgressCalc(isDone int, taskNum int) float64 {
	result := float64(isDone) / (float64(taskNum/15) * float64(isDone))
	if math.IsNaN(result) {
		result = 1
	}
	return result
}

// 一週間の進捗を計算
func toWeekProgressCalc(isDone int, taskNum int) float64 {
	result := float64(isDone) / float64(taskNum)
	if math.IsNaN(result) {
		result = 1
	}
	return result
}

func progressCalc(progress *progress) float64 {
	return progress.toDay*0.6 + progress.toWeek*0.4
}

func dbGetUserToken(userID string) (*Token, error) {
	res, err := http.Get(userReadByIdUrl + userID)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var userDataList domain.UserDataList
	if err := json.NewDecoder(res.Body).Decode(&userDataList); err != nil {
		return nil, err
	}
	user := userDataList[0]
	return &Token{
		AccessToken:  user.Token,
		RefreshToken: user.RefreshToken}, nil
}

func dbPostUser(user *domain.UserData) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	res, err := http.Post(userCreateUrl, "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var resMessage domain.UserDataSuccessMessage
	if err := json.NewDecoder(res.Body).Decode(&resMessage); err != nil {
		return err
	}
	log.Println(resMessage.Message)
	if !resMessage.Result {
		return errors.New("POST /user/create Faild")
	}
	return nil
}
