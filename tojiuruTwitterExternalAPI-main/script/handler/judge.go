package handler

import (
	"app/domain"
	"app/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

// POST /tojiuru/judge
func judge(c *gin.Context) {
	judge := domain.PostJudge{}
	err := c.Bind(&judge)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	userID := GetUserIDFromContext(c)
	if userID == "" {
		log.Println("Not set user id")
		c.JSON(400, "Not UserID")
		return
	}
	if err = usecase.CreateJudge(&usecase.Check{
		MessageId: judge.MessageId,
		UserId:    userID,
		Judge:     judge.Judge}); err != nil {
		c.JSON(500, "Internal Server error")
		return
	}
	c.JSON(201, "Create Success")
}
