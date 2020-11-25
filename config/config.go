package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

const (
	CONFIG_NAME =".env"
	CONFIG_TYPE = "yml"
	CONFIG_PATH = "./"
)

func Load(app *iris.Application) {
	v := viper.New()
	v.SetConfigName(CONFIG_NAME)
	v.SetConfigType(CONFIG_TYPE)
	v.AddConfigPath(CONFIG_PATH)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		app.Logger().Info("Config file changed: ", e.Name)
		app.Logger().Info(v.AllSettings())
	})

	err := v.ReadInConfig()
	if err != nil {
		app.Logger().Error(err)
	}

	configs := v.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
	app.Logger().Info(v.AllSettings())
}
