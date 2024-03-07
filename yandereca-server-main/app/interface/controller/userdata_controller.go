package controller

import (
	"fmt"
	"log"

	"yandereca.tech/yandereca/domain"
	"yandereca.tech/yandereca/interface/database"
	"yandereca.tech/yandereca/usecase"
)

type UserDataController struct {
	Interactor usecase.UserDataInteractor
}

func NewUserDataController(sqlHandler database.SqlHandler) *UserDataController {
	return &UserDataController{
		Interactor: usecase.UserDataInteractor{
			UserDataRepository: &database.UserDataRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserDataController) Create(c Context) {
	// テストオブジェクト

	userdata := domain.UserData{}
	err := c.Bind(&userdata)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(userdata.Id)
	u := domain.UserDataSuccessMessage{}

	user_id, err := controller.Interactor.Create(userdata)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	u.UserId = user_id.UserId
	u.Message = "success"
	u.Result = true
	c.JSON(201, u)
}

func (controller *UserDataController) TestCreate(c Context) {
	// テストオブジェクト
	userdata := domain.UserData{}
	userdata.Name = "example-name"
	userdata.Email = "example-email"
	userdata.Token = "example-token"
	userdata.RefreshToken = "example-refresh-token"
	userdata.GoogleUid = "example-google-uid"

	u := domain.UserDataSuccessMessage{}

	user_id, err := controller.Interactor.Create(userdata)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	u.UserId = user_id.UserId
	u.Message = "success"
	u.Result = true
	c.JSON(201, u)
}

func (controller *UserDataController) Read(c Context) {
	identifier := c.DefaultQuery("id", "example-id")
	userdata, err := controller.Interactor.UserDataRepository.FindById(identifier)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, userdata)
}

func (controller *UserDataController) Delete(c Context) {
	identifier := c.DefaultQuery("id", "example-id")
	err := controller.Interactor.UserDataRepository.Remove(identifier)

	u := domain.UserDataSuccessMessage{}

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	u.UserId = identifier
	u.Message = "success"
	u.Result = true
	c.JSON(201, u)
}
