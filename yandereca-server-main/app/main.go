package main

import (
	"yandereca.tech/yandereca/infra"
)

func main() {
	router := infra.NewRouter()
	router.Run()

	// router.GET("/hello", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello World")
	// })
	// router.Run(":8080")
}
