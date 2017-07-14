// +build appengine

package main

import (
	"github.com/wieringen/boilerplate/app/api"
	"net/http"
	"os"
)

func init() {
	os.Setenv("GAE_ENVIRONMENT", "true")

	router := api.GetRouter()

	http.Handle("/", router)
}
