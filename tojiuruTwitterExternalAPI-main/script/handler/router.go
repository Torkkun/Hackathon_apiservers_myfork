package handler

import (
	"app/handler/extwitter"
	"log"

	"github.com/dghubble/gologin/twitter"
	"github.com/dghubble/sessions"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

  "app/handler/env"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	config := extwitter.OAuthConfig()
	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			//"http://localhost:3000",
      env.ReadEnv("CBURL"),
		},
		//レスポンス公開の許可
		AllowCredentials: true,
	}))
	router.Use(authContext())
	// 静的ファイルサーバーとしてmediaファイル内のファイルを返す
	router.StaticFS("/image/", gin.Dir("../media/*", false))
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", extwitter.TwitterProfileHandler)
	router.POST("/tweet", gin.WrapH(extwitter.Tweethandler(config)))
	router.POST("/logout", logoutHandler())
	router.GET("/twitter/login", gin.WrapH(extwitter.LoginHandler(config, nil)))
	//twitter/loginの後呼び出される
	router.GET("/twitter/callback", gin.WrapH(twitter.CallbackHandler(config, extwitter.IssueSession(), nil)))
	router.GET("/check", gin.WrapH(extwitter.LoginCheck(config)))
	testRoute := router.Group("/test")
	{
		testRoute.GET("/mediatweet", func(c *gin.Context) {
			extwitter.TESTMEDIATWEET(c)
		})
		testRoute.POST("/teste", func(c *gin.Context) {
			testexamination(c)
		})
		testRoute.POST("/media", func(c *gin.Context) {
			testmedia(c)
		})
	}
	examinationRoute := router.Group("/examination")
	{
		examinationRoute.POST("/create", func(c *gin.Context) {
			examination(c)
		})
		examinationRoute.DELETE("/delete", func(c *gin.Context) {
			deleteData(c)
		})
		examinationRoute.GET("/data", func(c *gin.Context) {
			examinationData(c)
		})
	}
	judgeRoute := router.Group("/judge")
	{
		judgeRoute.POST("/create", func(c *gin.Context) {
			judge(c)
		})
	}
	mediaRoute := router.Group("/media")
	{
		mediaRoute.POST("/upload", func(c *gin.Context) {
			mediaUpload(c)
		})
	}
	replyRoute := router.Group("/reply")
	{
		replyRoute.GET("/users", func(c *gin.Context) {
			replyusers(c)
		})
		replyRoute.POST("/create", func(c *gin.Context) {
			replycreate(c)
		})
		replyRoute.GET("/data", func(c *gin.Context) {
			replydata(c)
		})
	}
	tojiuruRoute := router.Group("/tojiuru")
	{
		tojiuruRoute.POST("/authenticate", func(c *gin.Context) {
			authenticate(c)
		})
		tojiuruRoute.POST("/signup_account", func(c *gin.Context) {
			signupAccount(c)
		})
		tojiuruRoute.GET("/signout_account", func(c *gin.Context) {
			signoutAccount(c)
		})
		tojiuruRoute.GET("/data", func(c *gin.Context) {
			userdata(c)
		})
    //bot情報保存
    tojiuruRoute.POST("/link_service", func(c *gin.Context) {
      savefile_botinfo(c)
    })
	}
	//定期実行
	router.GET("/execution", func(c *gin.Context) {
		periodicExecution(c)
	})
  
	Router = router
}

// Context変換ginとの結びつき
func logoutHandler() gin.HandlerFunc {
	h := extwitter.LogoutHandler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func authContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// cookieからuseridを取得
		cookie, err := ctx.Request.Cookie("_cookie")
		if err != nil {
			log.Println(err)
			return
		}
		// コンテキストにuserIDを放り込む
		SetUserID(ctx, cookie.Value)

		// twitterのcookie
		var sessionStore = sessions.NewCookieStore([]byte("example cookie signing secret"), nil)
		session, err := sessionStore.Get(ctx.Request, "example-twtter-app")
		if err != nil {
			log.Println(err.Error())
			return
		}
		token := session.Values["accessUserAccessToken"].(string)
		secret := session.Values["twitterUserAccessSecret"].(string)
		// コンテキストにtwitterTokenIDを放り込む
		SetTwitterToken(ctx, &Token{AccessToken: token, SecretToken: secret})

		ctx.Next()
	}
}

const (
	userIDKey   string = "userID"
	accessToken string = "accessToken"
	secretToken string = "secretToken"
)

type Token struct {
	AccessToken string
	SecretToken string
}

// コンテキストにuseridを格納
func SetUserID(ctx *gin.Context, userID string) {
	ctx.Set(userIDKey, userID)
}

func SetTwitterToken(ctx *gin.Context, token *Token) {
	ctx.Set(accessToken, token.AccessToken)
	ctx.Set(secretToken, token.SecretToken)
}

func GetUserIDFromContext(ctx *gin.Context) string {
	var userID string
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(string)
	}
	return userID
}

func GetTwitterTokenFromContext(ctx *gin.Context) *Token {
	var accessToken string
	if ctx.Value(accessToken) != nil {
		accessToken = ctx.Value(accessToken).(string)
	}
	var secretToken string
	if ctx.Value(accessToken) != nil {
		secretToken = ctx.Value(secretToken).(string)
	}
	return &Token{AccessToken: accessToken, SecretToken: secretToken}
}
