package handler

import "net/http"

func Init() {
	http.HandleFunc("/sxdemo/ping", Ping)
	http.HandleFunc("/sxdemo/register", RegisterNode)
}
