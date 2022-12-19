package config

import (
	"fmt"
	"os"

	"github.com/GeekSpotlight/go-gin-example/model"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	loadConfigurations()
	// loadEnvConfigurations()
}

func loadConfigurations() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	viper.RegisterAlias("appName", "application.name")
}

func loadEnvConfigurations() {
	os.Setenv("port", "13")
	os.Setenv("TGC_PORT", "14")
	// viper.SetEnvPrefix("TGC")

	viper.BindEnv("port")
	// fmt.Printf("port '%v'\n", viper.Get("port"))

	// os.Setenv("TGC_PORT", "13") // typically done outside of the app
}

func Get(name string) interface{} {
	return viper.Get(name)
}

func GetAppConfig() model.AppConfig {
	var appConfig model.AppConfig

	viper.Unmarshal(&appConfig)

	return appConfig
}
