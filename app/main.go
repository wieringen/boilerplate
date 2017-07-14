package main

import (
	"github.com/wieringen/boilerplate/app/api"
	"net/http"
	"os"
)

func main() {
	os.Setenv("GAE_ENVIRONMENT", "false")

	router := api.GetRouter()

	http.ListenAndServe(":8080", router)
}
