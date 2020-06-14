package main

import (
	"net/http"

	"github.com/joincloud/home-platform/handler"
	"github.com/joincloud/home-platform/registry"
	"github.com/joincloud/home-platform/registry/memory"
)

func main() {
	registry.DefaultRegistry = &memory.Registry{}

	handler.Init()
	http.ListenAndServe(":8090", nil)
}
