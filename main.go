package main

import (
	"github.com/DomineCore/brisk"
	"github.com/DomineCore/briskdemo/internal/api/config"
	"github.com/DomineCore/briskdemo/internal/api/home"
	"github.com/DomineCore/briskdemo/internal/app"
)

func LoadUrl(b *brisk.Brisk) {
	b.Router.Use(&brisk.LoggingMiddleware{})
	b.Router.Include("/config/", &config.ConfigUrl)
	b.Router.Include("/home/", &home.HomeUrl)
}
func main() {
	app := app.GetApp()
	app.Conf.AddConfigPath("./config")
	app.Conf.SetConfigName("settings")
	app.Conf.SetConfigType("json")
	LoadUrl(app)
	app.Run(":8000")
}
