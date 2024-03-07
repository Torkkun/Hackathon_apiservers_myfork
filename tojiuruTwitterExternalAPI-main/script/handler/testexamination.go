package handler

import (
	"app/database"
	"app/domain"
	"log"

	"github.com/gin-gonic/gin"
)

func testexamination(c *gin.Context) {
	examination := domain.PostExamination{}
	err := c.Bind(&examination)
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
		return
	}
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
		return
	}
	cookie, err := c.Request.Cookie("_cookie")
	if err != nil {
		log.Println(err)
		c.JSON(500, err)
		return
	}
	//token := "test_token"
	//secret := "test_secret"
	uid := cookie.Value
	data := database.Examination{
		Message:  examination.Message,
		People:   examination.People,
		Deadline: examination.Deadline,
		//AccessToken: token,
		//SecretToken: secret,
		UserId: uid,
	}
	if err := database.CreateExamination(&data); err != nil {
		log.Println(err)
		c.JSON(500, err)
		return
	}
	c.JSON(201, "create_succese")
}
