package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GeekSpotlight/go-gin-example/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.New()

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

	// http://localhost:8080/hi
	router.GET("/hi", func(ctx *gin.Context) {
		appConfig := config.GetAppConfig()
		// appName := config.Get("appName")
		ctx.String(http.StatusOK, fmt.Sprintf("hello there! welcome to %s", appConfig.AppName))
	})

	// http://localhost:8080/hi
	router.GET("/hi", func(ctx *gin.Context) {
		appConfig := config.GetAppConfig()
		// appName := config.Get("appName")
		ctx.String(http.StatusOK, fmt.Sprintf("hello there! welcome to %s", appConfig.AppName))
	})

	// by default, runs in port :8080
	router.Run()
	// router.Run(":5001")
}

func main2() {
	// singleParam()
	// doubleParam()
	// multiParam()
	automatic()
}

func singleParam() {
	// setting up env variables
	os.Setenv("TGC_DATAPREFIX", "data with prefix TGC")
	os.Setenv("DATA_NO_PREFIX", "got no prefix")

	vPrefix := viper.New()
	vPrefix.SetEnvPrefix("TGC")
	vPrefix.BindEnv("dataprefix")

	vNoPrefix := viper.New()
	vNoPrefix.BindEnv("data_no_prefix")

	fmt.Println("output 1:", vPrefix.Get("dataprefix"))
	fmt.Println("output 2:", vNoPrefix.Get("data_no_prefix"))
}
func doubleParam() {
	// setting up env variables
	os.Setenv("TGC_DATA", "data with prefix TGC")
	os.Setenv("env_data", "got no prefix")

	vPrefix := viper.New()
	vPrefix.SetEnvPrefix("TGC")
	vPrefix.BindEnv("data", "env_data")        // does not append prefix and is not in full uppercase.
	vPrefix.BindEnv("data_prefix", "TGC_DATA") // explicitly mention full name of the param. prefix is ignored.

	fmt.Println("output 1:", vPrefix.Get("data"))
	fmt.Println("output 2:", vPrefix.Get("data_prefix"))
}

func multiParam() {
	// setting up env variables
	os.Setenv("high_precedence", "precedence level 1")
	os.Setenv("normal_precedence", "precedence level 2")
	os.Setenv("low_precedence", "precedence level 3")
	v := viper.New()

	v.BindEnv("precedence", "no_env_variable_configured", "high_precedence", "normal_precedence", "low_precedence")

	fmt.Println("output:", v.Get("precedence"))
}

func automatic() {
	// setting up env variables
	os.Setenv("TGC_AUTO_1", "an automatic env")
	os.Setenv("TGC_AUTO_2", "another automatic env")

	vPrefix := viper.New()
	vPrefix.SetEnvPrefix("TGC")
	vPrefix.AutomaticEnv()

	vNoPrefix := viper.New()
	vNoPrefix.AutomaticEnv()

	fmt.Println("output 1: ", vPrefix.Get("auto_1"))
	fmt.Println("output 2: ", vPrefix.Get("auto_2"))
	fmt.Println("output 3: ", vNoPrefix.Get("TGC_AUTO_1"))
	fmt.Println("output 4: ", vNoPrefix.Get("TGC_AUTO_2"))
}
