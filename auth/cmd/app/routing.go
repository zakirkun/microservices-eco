package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/microservices-eco/auth/pkg/config"
)

func init() {

}

func Router(debug bool) http.Handler {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	app := config.NewGin()
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Oke")
	})

	var version = "/v1/"
	v1 := app.Group(version)
	v1.GET("auth/login", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	return app
}
