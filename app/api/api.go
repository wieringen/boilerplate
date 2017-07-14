package api

import (
	"net/http"
	"io"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func GetRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	return mux
}
