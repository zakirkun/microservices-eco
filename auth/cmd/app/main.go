package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/microservices-eco/auth/database/migrations"
	"github.com/zakirkun/microservices-eco/auth/pkg/config"
)

type ParsedFlags struct {
	serverAddr string
	serverPort int
}

var pf ParsedFlags

func init() {
	flag.StringVar(&pf.serverAddr, "serverAddr", "", "HTTP server network address")
	flag.IntVar(&pf.serverPort, "serverPort", 4000, "HTTP server network port")
	flag.Parse()
}

func main() {
	router := config.NewGin()
	configuration := config.New()
	db := config.NewDatabase(configuration)

	migration := migrations.New(db)
	_ = migration.Seeder()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Oke")
	})

	serverURI := fmt.Sprintf("%s:%d", pf.serverAddr, pf.serverPort)
	s := &http.Server{
		Addr:           serverURI,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Running Server
	s.ListenAndServe()
}
