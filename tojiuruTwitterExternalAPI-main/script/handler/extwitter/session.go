package extwitter

import (
	"app/handler/env"
	"log"
	"net/http"

	"github.com/dghubble/gologin/oauth1"
	"github.com/dghubble/gologin/twitter"

	"github.com/dghubble/sessions"
)

const (
	sessionName             = "example-twtter-app"
	sessisonSecret          = "example cookie signing secret"
	sessionUserKey          = "twitterID"
	sessionUsername         = "twitterUsername"
	sessionUserAccessToken  = "accessUserAccessToken"
	sessionUserAccessSecret = "twitterUserAccessSecret"
)

var sessionStore = sessions.NewCookieStore([]byte(sessisonSecret), nil)

// セッション
func IssueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		twitterUser, err := twitter.UserFromContext(ctx)
		//fmt.Println(twitterUser)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 2. Implement a success handler to issue some form of session
		accessToken, accessSecret, err := oauth1.AccessTokenFromContext(ctx)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session := sessionStore.New(sessionName)
		session.Values[sessionUserKey] = twitterUser.ID
		session.Values[sessionUsername] = twitterUser.ScreenName
		session.Values[sessionUserAccessToken] = accessToken
		session.Values[sessionUserAccessSecret] = accessSecret
		session.Save(w)

		http.Redirect(w, req, env.ReadEnv("CBURL"), http.StatusFound)

	}
	return http.HandlerFunc(fn)

}
