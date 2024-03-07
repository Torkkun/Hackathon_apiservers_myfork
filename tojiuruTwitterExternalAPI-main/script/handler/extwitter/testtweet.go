package extwitter

import (
	"app/domain"
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var testmedia = Test

func testclient(accessToken string, accessSecret string) *anaconda.TwitterApi {
	err := godotenv.Load(".env")
	if err != nil {
		// .env読めなかった場合の処理
		log.Fatalln(err)
	}
	anaconda.SetConsumerKey(os.Getenv("Twitter_Consumer"))
	anaconda.SetConsumerSecret(os.Getenv("Twitter_Secret"))
	client := anaconda.NewTwitterApi(accessToken, accessSecret)
	return client
}

func medeaUpload(tweet *domain.Tweet) (mediaid int64, err error) {
	client := testclient(tweet.AccessToken, tweet.SecretToken)
	media, err := client.UploadMedia(tweet.Media)
	mediaid = media.MediaID
	return
}

func test_tweet(tweet *domain.Tweet) error {
	twitterClient := getClient(tweet.AccessToken, tweet.SecretToken)
	param := new(twitter.StatusUpdateParams)
	param.MediaIds = tweet.MediaID
	_, _, err := twitterClient.Statuses.Update(tweet.Tweet_text, param)
	if err != nil {
		return err
	}
	fmt.Println("uppdated tweet")
	return err
}

func TESTMEDIATWEET(c *gin.Context) {
	var sessionStore = sessions.NewCookieStore([]byte("example cookie signing secret"), nil)
	session, err := sessionStore.Get(c.Request, "example-twtter-app")
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	token := session.Values["accessUserAccessToken"].(string)
	secret := session.Values["twitterUserAccessSecret"].(string)

	tw := &domain.Tweet{
		Tweet_text:  "hello",
		Media:       testmedia,
		AccessToken: token,
		SecretToken: secret,
	}
	media, err := medeaUpload(tw)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	var ml [4]int64 = [4]int64{media, media, media, media}
	//tw.MediaID = append(tw.MediaID, media)
	tw.MediaID = ml[:]
	test_tweet(tw)

}
