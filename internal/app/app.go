package app

import (
	"sync"

	"github.com/DomineCore/brisk"
)

var once sync.Once
var App *brisk.Brisk

func GetApp() *brisk.Brisk {
	once.Do(func() {
		App = brisk.New()
		// return instance
	})
	return App
}
