package home

import (
	"net/http"

	"github.com/DomineCore/brisk"
)

var HomeUrl brisk.Router = *brisk.NewRouter()

func init() {
	HomeUrl.Add("/one/", http.MethodGet, Home)
}
