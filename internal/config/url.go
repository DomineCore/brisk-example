package config

import (
	"net/http"

	"github.com/DomineCore/brisk"
)

var ConfigUrl brisk.Router = *brisk.NewRouter()

func init() {
	ConfigUrl.Add("/get/", http.MethodGet, Config)
}
