package googletodo

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

type Repository struct {
	Config *oauth2.Config
}

const (
	googleTasks           = "https://www.googleapis.com/auth/tasks"
	googleUsetInfoEmail   = "https://www.googleapis.com/auth/userinfo.email"
	googleUserInfoProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

// 他のAPIライブラリ等の様子を見て後々infra層に切り分けたい部分
func NewGoogleService() *Repository {
	// config setting
	jsonKey, err := ioutil.ReadFile("app/credentials.json")
	if err != nil {
		// ファイルを読み込めてない場合Exit(1)を呼び出し終了
		log.Fatalf("could not open the file for google tasks settings: %s", err)
	}
	scopes := []string{
		googleTasks,
		googleUsetInfoEmail,
		googleUserInfoProfile}
	config, err := google.ConfigFromJSON(jsonKey, scopes...)
	return &Repository{
		Config: config}
}

// configから取得
func (repo *Repository) GoogleURL() string {
	authURL := repo.Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return authURL
}

// configExchangeで取得
func (repo *Repository) GoogleToken(authCode string) (*oauth2.Token, error) {
	token, err := repo.Config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}
	return token, nil
}

type UserInfo struct {
	Id         string
	Name       string
	Email      string
	FamilyName string
	GivenName  string
}

type claims struct {
	Subject       string `json:"sub"`
	Name          string `json:"name"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

// tokenを使用し取得
func (repo *Repository) GoogleUserInfo(token *oauth2.Token) (*UserInfo, error) {
	tokenSource := repo.Config.TokenSource(context.Background(), token)
	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return nil, err
	}
	userInfo, err := provider.UserInfo(context.Background(), tokenSource)
	if err != nil {
		return nil, err
	}
	cl := new(claims)
	if err := userInfo.Claims(cl); err != nil {
		return nil, err
	}
	return &UserInfo{
		Id:         cl.Subject,
		Name:       cl.Name,
		FamilyName: cl.FamilyName,
		GivenName:  cl.GivenName,
		Email:      cl.Email}, nil
}

type GoogleTaskService struct {
	Srv *tasks.Service
}

func (repo *Repository) NewGoogleTaskService(token *oauth2.Token) (*GoogleTaskService, error) {
	client := repo.Config.Client(context.Background(), token)
	tasksrv, err := tasks.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	return &GoogleTaskService{Srv: tasksrv}, nil
}

// ※任意の値を設定できるようにするdayとweekで分けない

// tasksAPIで取得
func (ggtask *GoogleTaskService) GoogleTaskLists() (*tasks.TaskLists, error) {
	return ggtask.Srv.Tasklists.List().Do()
}

type GoogleTaskRequest struct {
	ListId []string
	DueMax string
	DueMin string
}

type GoogleTaskResponse struct {
	Tasks     []*tasks.Task
	IsDoneNum int
}

func (ggtask *GoogleTaskService) GoogleTasks(ggtrq *GoogleTaskRequest) (*GoogleTaskResponse, error) {
	var allTasks []*tasks.Task
	var isDone int
	for _, id := range ggtrq.ListId {
		// GoogleTasksAPIにGetリクエストを送る
		tasks, err := ggtask.Srv.Tasks.List(id).DueMax(ggtrq.DueMax).DueMin(ggtrq.DueMin).ShowCompleted(true).ShowHidden(true).Do()
		if err != nil {
			return nil, err
		}
		// hiddenパラメーターでタスクが完了しているかどうかを確かめる
		for _, i := range tasks.Items {
			if i.Hidden {
				isDone += 1
			}
		}
		allTasks = append(allTasks, tasks.Items...)
	}
	return &GoogleTaskResponse{
		Tasks:     allTasks,
		IsDoneNum: isDone}, nil
}
