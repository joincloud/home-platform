package main

import (
	"net/http"

	"github.com/joincloud/home-platform/handler"
)

func main() {
	handler.Init()
	http.ListenAndServe(":8090", nil)
}
