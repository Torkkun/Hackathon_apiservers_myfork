package extwitter

import (
	"app/domain"
	"app/handler/env"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	oauth1L "github.com/dghubble/gologin/oauth1"
	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type tweetcontent struct {
	Text  string `json:"text"`
	Tweet string `json:"tweet"`
}

// tweet
func Tweethandler(config *oauth1.Config) http.Handler {
	tweet := func(w http.ResponseWriter, req *http.Request) {
		//ここでやる場合,コールバックされたときに再度tweetする必要がある。
		//セッションの確認
		_, err := sessionStore.Get(req, sessionName)
		if err != nil {
			fmt.Println("twitter認証がされてない場合ログイン")
			//twitter認証がされてない場合ログイン
			log.Print(err)
			http.Redirect(w, req, "/twitter/login", http.StatusFound)

		}
		var v tweetcontent
		json.NewDecoder(req.Body).Decode(&v)
		fmt.Println(v)
		var message string = v.Text

		twitterClient, err := getTwitterClientFromRequest(config, req)
		if err != nil {
			log.Print(err)
			return
		}

		_, _, err = twitterClient.Statuses.Update(message, nil)
		if err != nil {
			log.Print(err)
			return
		}
		http.Redirect(w, req, env.ReadEnv("CBURL"), http.StatusFound)
	}
	return http.HandlerFunc(tweet)

}

func getTwitterClientFromRequest(config *oauth1.Config, req *http.Request) (*twitter.Client, error) {
	session, err := sessionStore.Get(req, sessionName)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	accessToken := session.Values[sessionUserAccessToken].(string)
	accessSecret := session.Values[sessionUserAccessSecret].(string)
	httpClient := config.Client(req.Context(), oauth1.NewToken(accessToken, accessSecret))
	twitterClient := twitter.NewClient(httpClient)
	return twitterClient, nil
}

// ツイートする用newTweetfunction
func Tweet(tweet domain.Tweet) error {
	twitterClient := getClient(tweet.AccessToken, tweet.SecretToken)
	_, _, err := twitterClient.Statuses.Update(tweet.Tweet_text, nil)
	if err != nil {
		return err
	}
	fmt.Println("uppdated tweet")
	return err
}

func getClient(accessToken string, accessSecret string) *twitter.Client {
	err := godotenv.Load(".env")
	if err != nil {
		// .env読めなかった場合の処理
		log.Fatalln(err)
	}

	consumerconfig := &TwitterConfig{
		ConsumerKey:    os.Getenv("Twitter_Consumer"),
		ConsumerSecret: os.Getenv("Twitter_Secret"),
	}
	config := oauth1.NewConfig(consumerconfig.ConsumerKey, consumerconfig.ConsumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)
	return twitterClient
}

// logout
func LogoutHandler() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			sessionStore.Destroy(w, sessionName)
		}
		http.Redirect(w, req, "/", http.StatusFound)

	}
	return http.HandlerFunc(fn)

}

func TwitterProfileHandler(c *gin.Context) {
	//認証確認
	//_, err := sessionStore.Get(c.Request, sessionName)
	//if err != nil {
	//twitter認証していないときはhome.htmlを表示
	//log.Print(err)
	//c.HTML(200, "home.html", nil)
	//http.Redirect(c.Writer, c.Request, "/twitter/login", http.StatusFound)
	//return

	//}
	//twitter認証済みであればprofile.htmlを表示
	c.HTML(200, "profile.html", nil)
}

func LoginHandler(config *oauth1.Config, failure http.Handler) http.Handler {
	success := AuthRedirectHandler(config, failure)
	return OauthLoginHandler(config, success, failure)
}

func AuthRedirectHandler(config *oauth1.Config, failure http.Handler) http.Handler {
	if failure == nil {
		failure = gologin.DefaultFailureHandler
	}
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		//requestToken, _, err := gologin.RequestTokenFromContext(ctx)
		requestToken, _, err := oauth1L.RequestTokenFromContext(ctx)
		if err != nil {
			fmt.Println(err)
			ctx = gologin.WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))
			return
		}
		authorizationURL, err := config.AuthorizationURL(requestToken)
		fmt.Println(authorizationURL.String())
		if err != nil {
			fmt.Println(err)
			ctx = gologin.WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))
			return
		}

		defer func() {
			url, e := json.Marshal(authorizationURL.String())
			if e != nil {
				fmt.Println(e)
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(url))
		}()
		//ここでtwitter社の認証画面にリダイレクトされる
		//http.Redirect(w, req, authorizationURL.String(), http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func OauthLoginHandler(config *oauth1.Config, success, failure http.Handler) http.Handler {
	if failure == nil {
		failure = gologin.DefaultFailureHandler
	}
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		requestToken, requestSecret, err := config.RequestToken()
		if err != nil {
			ctx = gologin.WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))
			return
		}
		ctx = oauth1L.WithRequestToken(ctx, requestToken, requestSecret)
		success.ServeHTTP(w, req.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func LoginCheck(config *oauth1.Config) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		_, err := sessionStore.Get(req, sessionName)
		if err != nil {
			log.Print(err)

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, "no-login")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "login")
	}
	return http.HandlerFunc(fn)
}
