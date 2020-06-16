package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/joincloud/home-platform/handler"
	"github.com/joincloud/home-platform/registry"
	"github.com/joincloud/home-platform/registry/memory"
)

func main() {
	registry.DefaultRegistry = &memory.Registry{}

	router := mux.NewRouter()
	router.HandleFunc("/sxdemo/ping", handler.Ping)
	router.HandleFunc("/sxdemo/register", handler.RegisterNode)
	http.Handle("/sxdemo", router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":8090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
