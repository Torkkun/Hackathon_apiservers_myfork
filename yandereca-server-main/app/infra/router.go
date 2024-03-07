package infra

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"yandereca.tech/yandereca/interface/controller"
	"yandereca.tech/yandereca/interface/controller/todo"
	"yandereca.tech/yandereca/interface/googletodo"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	router := &Router{
		Engine: gin.Default(),
	}
	handler := NewEnvHandler()
	url, _ := handler.ReadEnv("CORS_URL")
	sqlHandler := NewSqlHandler()
	googleService := googletodo.NewGoogleService()

	router.Engine.Use(cors.New(cors.Config{
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
			url,
			"http://localhost:3000",
		},
		MaxAge: 24 * time.Hour,
	}))

	//コントローラー単位でルーティング関数を定義
	//コントローラー作成
	taskController := controller.NewTaskController(sqlHandler)
	userdataController := controller.NewUserDataController(sqlHandler)
	googleToDoController := todo.NewGoogleToDoController(googleService)

	//ルーティングテーブルの初期化
	router.initTaskRoute(*taskController)
	router.initUserDataRoute(*userdataController)
	router.initGoogleToDoRoute(*googleToDoController)

	return router
}

func (router *Router) initTaskRoute(taskController controller.TaskController) {
	taskRoute := router.Engine.Group("/task")
	{
		taskRoute.POST("/create", func(c *gin.Context) {
			taskController.Create(c)
		})
		taskRoute.POST("/update", func(c *gin.Context) {
			taskController.Create(c)
		})
		taskRoute.POST("/read", func(c *gin.Context) {
			taskController.Read(c)
		})
		taskRoute.GET("/read", func(c *gin.Context) {
			taskController.Read(c)
		})
		taskRoute.GET("/progress", func(c *gin.Context) {
			taskController.Calc(c)
		})
		testTaskRoute := taskRoute.Group("/test")
		{
			testTaskRoute.GET("/create", func(c *gin.Context) {
				taskController.TestCreate(c)
			})
			testTaskRoute.GET("/read", func(c *gin.Context) {
				taskController.Read(c)
			})
		}
	}
}

func (router *Router) initUserDataRoute(userdataController controller.UserDataController) {
	userdataRoute := router.Engine.Group("/user")
	{
		userdataRoute.POST("/create", func(c *gin.Context) {
			userdataController.Create(c)
		})
		userdataRoute.POST("/update", func(c *gin.Context) {
			userdataController.Create(c)
		})
		userdataRoute.POST("/read", func(c *gin.Context) {
			userdataController.Read(c)
		})
		userdataRoute.GET("/read", func(c *gin.Context) {
			userdataController.Read(c)
		})
		userdataRoute.DELETE("/delete", func(c *gin.Context) {
			userdataController.Delete(c)
		})
		testUserdataRoute := userdataRoute.Group("/test")
		{
			testUserdataRoute.GET("/create", func(c *gin.Context) {
				userdataController.TestCreate(c)
			})
			testUserdataRoute.GET("/read", func(c *gin.Context) {
				userdataController.Read(c)
			})
		}
	}
}

//外部API
func (router *Router) initGoogleToDoRoute(ggCon todo.GoogleController) {
	googleTaskRoute := router.Engine.Group("/googletask")
	{
		googleTaskRoute.GET("/request", func(c *gin.Context) {
			ggCon.GetURL(c)
		})

		googleTaskRoute.POST("/auth", func(c *gin.Context) {
			ggCon.PostAuthCode(c)
		})

		googleTaskRoute.GET("/progress", func(c *gin.Context) {
			ggCon.GetProgress(c)
		})

	}
}

func (router *Router) initTestRoute() {
	router.Engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
}

func (router *Router) Run() {
	router.Engine.Run()
}
