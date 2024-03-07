package todo

import (
	"errors"
	"time"

	"golang.org/x/oauth2"
	"yandereca.tech/yandereca/domain"
	"yandereca.tech/yandereca/interface/googletodo"
)

type GoogleInteractor struct {
	GoogleRepo GoogleRepository
}

type GoogleRepository interface {
	GoogleURL() string
	GoogleToken(authCode string) (*oauth2.Token, error)
	GoogleUserInfo(*oauth2.Token) (*googletodo.UserInfo, error)
	NewGoogleTaskService(*oauth2.Token) (*googletodo.GoogleTaskService, error)
}

// AuthURLを生成する
func (gi *GoogleInteractor) CreateURL() string {
	return gi.GoogleRepo.GoogleURL()
}

// AuthCodeをTokenに変換する
func (gi *GoogleInteractor) AuthCode(authCode string) (*Token, error) {
	token, err := gi.GoogleRepo.GoogleToken(authCode)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken}, nil
}

// googleからuserinfo情報を取得し返還するJwtを生成
func (gi *GoogleInteractor) CreateJwt(token *Token) (*Signed, error) {
	// googleのuserinfoを取得
	googleUserInfo, err := gi.GoogleRepo.GoogleUserInfo(&oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    "Bearer"})
	if err != nil {
		return nil, err
	}
	signedToken, err := newJwtCreate(&userJson{
		UserId:   googleUserInfo.Id,
		FullName: googleUserInfo.Name,
		SurName:  googleUserInfo.FamilyName,
		Name:     googleUserInfo.GivenName,
	})
	if err != nil {
		return nil, err
	} else if signedToken == "" {
		return nil, errors.New("jwt could not be created successfully")
	}
	// 最後にuser/createにPOSTする
	if err := dbPostUser(&domain.UserData{
		Id:           googleUserInfo.Id,
		Name:         googleUserInfo.Name,
		Email:        googleUserInfo.Email,
		Token:        token.AccessToken,
		RefreshToken: token.RefreshToken,
		GoogleUid:    googleUserInfo.Id}); err != nil {
		return nil, err
	}
	return &Signed{Jwt: signedToken}, nil
}

// 進捗率を計算する
func (gi *GoogleInteractor) CalcProgress(userId string) (*Progress, error) {
	// tokenを取得
	// 今のところ自APIアクセスする形で取得
	token, err := dbGetUserToken(userId)
	if err != nil {
		return nil, err
	}
	// リストの取得とタスクの取得をするサービス
	srv, err := gi.GoogleRepo.NewGoogleTaskService(&oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    "Bearer",
	})
	if err != nil {
		return nil, err
	}
	// googletodoのtasklistをごと持ってくる
	tasklists, err := srv.GoogleTaskLists()
	if err != nil {
		return nil, err
	}
	// IDだけ取り出す
	var idlist []string
	for _, i := range tasklists.Items {
		idlist = append(idlist, i.Id)
	}
	nowtime := time.Now()
	// 最小(今日の日付)
	duemin := formatTime(time.Date(nowtime.Year(), nowtime.Month(), nowtime.Day(), 0, 0, 0, 0, time.Local))
	// tasklistごとのタスクのデータを取得
	// 今日のタスク
	day, err := srv.GoogleTasks(&googletodo.GoogleTaskRequest{
		ListId: idlist,
		DueMax: formatTime(time.Date(nowtime.Year(), nowtime.Month(), nowtime.Day()+1, 0, 0, 0, 0, time.Local)),
		DueMin: duemin})
	// 今日から一週間のタスク
	week, err := srv.GoogleTasks(&googletodo.GoogleTaskRequest{
		ListId: idlist,
		DueMax: formatTime(time.Date(nowtime.Year(), nowtime.Month(), nowtime.Day()+8, 0, 0, 0, 0, time.Local)),
		DueMin: duemin})
	// 進捗を計算し返す
	result := progressCalc(&progress{
		toDay:  toDayProgressCalc(day.IsDoneNum, len(day.Tasks)),
		toWeek: toWeekProgressCalc(week.IsDoneNum, len(week.Tasks))})
	return &Progress{Result: result}, nil
}
