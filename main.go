package main

import (
	"github.com/DomineCore/brisk"
	"github.com/DomineCore/briskdemo/internal/config"
	"github.com/DomineCore/briskdemo/internal/home"
)

func LoadUrl(b *brisk.Brisk) {
	b.Router.Use(&brisk.LoggingMiddleware{})
	b.Router.Include("/config/", &config.ConfigUrl)
	b.Router.Include("/home/", &home.HomeUrl)
}

func MigrateDB() {
	brisk.DB.AutoMigrate(
		&config.Model{},
	)
}

func main() {
	app := brisk.New("./config/settings.json")
	LoadUrl(app)
	MigrateDB()
	app.Run(":8000")
}
