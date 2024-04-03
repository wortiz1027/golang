package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

var (
	server *gin.Engine
)

func init() {
	server = gin.Default()
}

func main() {

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	logger.Info().Msg("Configuring app...")

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:8080"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/")
	router.GET("healthcheck", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})

	server.Run(":3000")
}
