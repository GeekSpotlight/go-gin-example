package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// serves http://localhost:8080/health
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "I'm alive!")
	})

	// http://localhost:8080/hi
	router.GET("/hi", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello there!")
	})

	// http://localhost:8080/bye
	router.GET("/bye", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "take care now")
	})

	// by default, runs in port :8080
	router.Run()
	// router.Run(":5001")
}
