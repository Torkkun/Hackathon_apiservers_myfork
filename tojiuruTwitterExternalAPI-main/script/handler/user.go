package handler

import (
	"app/database"
	"log"

	"github.com/gin-gonic/gin"
)

// GET /tojiuru/userdata
func userdata(c *gin.Context) {
	users, err := database.GetUsers()
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
		return
	}
	// エンティティがない
	c.JSON(201, users)
}
