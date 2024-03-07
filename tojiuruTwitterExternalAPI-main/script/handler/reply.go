package handler

import (
	"app/database"
	"app/domain"
	"app/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

// GET /reply/users
func replyusers(c *gin.Context) {
	message_id := c.Query("message_id")
	if message_id == "" {
		c.JSON(400, "Bad Request")
		return
	}
	fromUserId := GetUserIDFromContext(c)
	// examinationTableから照合
	isOther, err := usecase.CheckUserIdFromExamination(message_id, fromUserId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, "Internal Server Error")
		return
	}
	if isOther {
		if err := usecase.CreateReplyUser(message_id, fromUserId); err != nil {
			log.Println(err.Error())
			c.JSON(500, "Internal Server Error")
		}
	}
	// replyUserTableからreplyIDを取得
	replyUserList, err := usecase.GetReplyUser(message_id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, "Internal Server Error")
	}
	c.JSON(201, replyUserList)
}

// POST /reply/create
// replyIDはReplyUserテーブルで作成されるやつ
// チェックしてないのでIDがわかれば誰でもメッセージ書き込めるようになってる。だけど何すればいいのか分からない
func replycreate(c *gin.Context) {
	reply := domain.PostReply{}
	err := c.Bind(&reply)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, "Internal Server Error")
		return
	}
	fromuid := GetUserIDFromContext(c)
	if err := usecase.CreateReplyMessage(reply.ReplyID, fromuid, reply.ReplyText); err != nil {
		log.Println(err.Error())
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(201, "create_succese")
}

// GET /reply/data
func replydata(c *gin.Context) {
	reply_id := c.Query("reply_id")
	if reply_id == "" {
		log.Println("no reply_id")
		c.JSON(400, "Bad Request")
		return
	}
	ruserid := GetUserIDFromContext(c)
	// userをチェックする
	if err := usecase.CheckUserIdFromReplyUser(ruserid, reply_id); err != nil {
		log.Println(err.Error())
		c.JSON(500, "Internal Server Error")
		return
	}
	data, err := database.GetReplyMessage(reply_id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(200, data)
}
