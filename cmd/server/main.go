//go:generate go run github.com/google/wire/cmd/wire
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hum2/backend/internal/interface/controller/user/gen"
	"github.com/hum2/backend/internal/interface/middleware"
)

func test() {

}
func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// TODO: middleware（認証サンプル）
	r.Use(middleware.SampleMiddleware())
	// CORS
	r.Use(middleware.CorsMiddleware(app.Conf))
	// logger
	r.Use(middleware.LoggerMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello, world",
		})
	})
	// /users, /user, /user/:id
	gen.RegisterHandlers(r, app.UserController)
	r.Run()
}
