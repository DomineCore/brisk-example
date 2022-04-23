package config

import (
	"net/http"

	"github.com/DomineCore/brisk"
)

type Database struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Config(c *brisk.Context) {
	username := brisk.Config.GetString("Databases.default.username")
	db := &Database{
		Username: username,
		Password: brisk.Config.GetString("Databases.default.password"),
	}
	c.WriteJSON(http.StatusOK, db)
}
