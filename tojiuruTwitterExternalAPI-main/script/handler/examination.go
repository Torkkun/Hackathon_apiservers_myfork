package handler

import (
	"app/database"
	"app/domain"
	"app/handler/extwitter"
	"app/usecase"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

// POST /examination/create
func examination(c *gin.Context) {
	examination := domain.PostExamination{}
	err := c.Bind(&examination)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, "Incorrect JSON")
		return
	}
	userID := GetUserIDFromContext(c)
	if userID == "" {
		log.Println("Not set user id")
		c.JSON(400, "Not UserID")
		return
	}
	token := GetTwitterTokenFromContext(c)
	if token == nil {
		log.Println("Not set Twitter Token")
		c.JSON(400, "Not Twitter login")
		return
	}
	if err := usecase.CreateExamination(c, userID, examination); err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(201, "create_succese")
}

// 現在使用されてない(ビジネスロジックを考え直す)
/* func to_tweet_or_not_to_tweet(message_id string) {
	edata, err := database.ExaminationByMessageID(message_id)
	if err != nil {
		log.Println(err)
		return
	}
	judgemap, err := database.CountJudge(message_id)
	if err != nil {
		log.Println(err)
		return
	}
	truenum := judgemap[true]
	if edata.People <= truenum {
		tweet := domain.Tweet{
			Tweet_text: edata.Message,
			//AccessToken: edata.AccessToken,
			//SecretToken: edata.SecretToken,
		}
		err = extwitter.Tweet(tweet)
		if err != nil {
			log.Println(err)
			return
		}
	}
} */

// GET /examination/data
func examinationData(c *gin.Context) {
	// userIDを取得
	userID := GetUserIDFromContext(c)
	examination, err := usecase.GetExamination(userID)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server error")
		return
	}
	c.JSON(201, examination)
}

// DELETE /examination/delete
func deleteData(c *gin.Context) {
	tweet_id := c.Query("tweet_id")
	err := database.DeleteExamination(tweet_id)
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, "success delete")
}

// 定期実行
// GET /execution
func periodicExecution(c *gin.Context) {
	tweetlist, err := database.Checkdeadline()
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no expiration date")
			return
		} else {
			log.Fatalln(err)
			c.JSON(501, err.Error())
		}
	} else {
		for _, i := range tweetlist {
			tweet := domain.Tweet{
				Tweet_text:  i.Tweet_text,
				AccessToken: i.AccessToken,
				SecretToken: i.SecretToken,
			}
			//ツイートする
			err = extwitter.Tweet(tweet)
			if err != nil {
				log.Println(err)
				c.JSON(500, err.Error())
				return
			}
			err = database.DeleteExamination(i.Tweet_ID)
			if err != nil {
				log.Println(err)
				c.JSON(500, err.Error())
				return
			}
		}
		c.JSON(200, "tweeted")
		return
	}
}
