package config

import (
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewGin() *gin.Engine {
	// init Gin Engine
	app := gin.Default()

	// make config cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	// init zap logging
	logger, _ := zap.NewDevelopment()

	/**
	* Setup Cors
	 */
	app.Use(cors.New(config))

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	app.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	app.Use(ginzap.RecoveryWithZap(logger, true))

	app.Use(gin.Recovery())

	return app
}
