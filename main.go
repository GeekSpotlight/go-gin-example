package main

import (
	"fmt"
	"net/http"

	"github.com/GeekSpotlight/go-gin-example/config"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	// router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gzip.Gzip(gzip.BestCompression))

	// serves http://localhost:8080/health
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "I'm alive!")
	})

	// http://localhost:8080/hi
	router.GET("/hi", func(ctx *gin.Context) {
		appName := config.Get("application.name")
		ctx.String(http.StatusOK, fmt.Sprintf("hello there! welcome to %s", appName))
	})

	// http://localhost:8080/bye
	router.GET("/bye", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "take care now")
	})

	// http://localhost:8080/large-data-static
	router.StaticFile("/large-data-static", "./resources/large_json_data.json")

	// by default, runs in port :8080
	router.Run()
	// router.Run(":5001")
}
