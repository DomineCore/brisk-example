package config

import (
	"net/http"

	"github.com/DomineCore/brisk"
	"github.com/DomineCore/briskdemo/internal/app"
)

type Database struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Config(c *brisk.Context) {
	myapp := app.GetApp()
	username := myapp.Conf.GetString("DB.default.username")
	db := &Database{
		Username: username,
		Password: myapp.Conf.GetString("DB.default.password"),
	}
	c.WriteJSON(http.StatusOK, db)
}
