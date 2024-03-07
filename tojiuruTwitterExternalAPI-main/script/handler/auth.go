package handler

import (
	"app/database"
	"app/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//POST /signup_account
func signupAccount(c *gin.Context) {
	var user domain.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		fmt.Printf("miss decode json :%v", err)
		c.JSON(500, err)
		return
	}
	if err = database.CreateUser(
		&database.User{
			Name:     user.Name,
			Password: user.Password,
		}); err != nil {
		fmt.Println(err)
		c.JSON(500, "This name is not available")
		return
	}
	c.JSON(200, "ok")
}

//POST /authenticate
func authenticate(c *gin.Context) {
	var user domain.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	u, err := database.UserByName(user.Name)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	// パスワード確かめ
	if u.Password == database.Encrypt(user.Password) {
		err := database.CreateSession(user.Name)
		if err != nil {
			log.Println(err)
			c.JSON(500, "Internal Server Error")
			return
		}
		// Samesite Sucureがどうか
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    u.UserId,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().AddDate(1, 0, 0),
		}
		http.SetCookie(c.Writer, &cookie)
		c.JSON(200, "login success")
	} else {
		c.JSON(403, "invalid password")
	}

}

//GET /signout_account
func signoutAccount(c *gin.Context) {
	userID := GetUserIDFromContext(c)
	if userID == "" {
		log.Println("Not set user id")
		c.JSON(400, "Not UserID")
		return
	}
	database.DeleteByUserID(&database.Session{UserId: userID})
	c.JSON(200, "logout")
}
