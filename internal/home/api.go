package home

import (
	"net/http"

	"github.com/DomineCore/brisk"
)

func Home(c *brisk.Context) {
	html := "<h1>标题一</h1>"
	c.HTML(http.StatusOK, html)
}
