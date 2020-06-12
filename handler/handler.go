package handler

import "net/http"

func Init() {
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/register", RegisterNode)
}
