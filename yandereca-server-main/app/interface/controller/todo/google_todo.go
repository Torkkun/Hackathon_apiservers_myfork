package todo

import (
	"log"

	"yandereca.tech/yandereca/domain"
	"yandereca.tech/yandereca/interface/controller"
	"yandereca.tech/yandereca/interface/googletodo"
	"yandereca.tech/yandereca/usecase/todo"
)

type GoogleController struct {
	ToDoInteractor todo.GoogleInteractor
}

// Injection
func NewGoogleToDoController(repository *googletodo.Repository) *GoogleController {
	return &GoogleController{
		ToDoInteractor: todo.GoogleInteractor{
			GoogleRepo: repository},
	}
}

// GET /googletask/request
func (con *GoogleController) GetURL(c controller.Context) {
	url := con.ToDoInteractor.CreateURL()
	c.JSON(200, domain.UrlResponse{AuthUrl: url})
}

// POST /googletask/auth
func (con *GoogleController) PostAuthCode(c controller.Context) {
	auth := domain.PostAuthCodeRequest{}
	err := c.Bind(&auth)
	if err != nil {
		log.Println(err)
		c.JSON(400, "Bad Request")
		return
	}
	token, err := con.ToDoInteractor.AuthCode(auth.Code)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	signed, err := con.ToDoInteractor.CreateJwt(&todo.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken})
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(200, signed.Jwt)
}

// GET /googletask/progress
func (con *GoogleController) GetProgress(c controller.Context) {
	userId := c.Query("uid")
	progress, err := con.ToDoInteractor.CalcProgress(userId)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(201, domain.CalcProgressResponse{Progress: progress.Result})
}
