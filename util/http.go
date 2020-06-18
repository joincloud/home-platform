package util

import (
	"encoding/json"
	"net/http"
)

type RspData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func RspJsonOK(w http.ResponseWriter, data interface{}) {
	rspData, _ := json.Marshal(RspData{
		Data:   data,
		Status: 200,
	})

	_, _ = w.Write(rspData)
}
