package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/joincloud/home-platform/registry"
)

func RegisterNode(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// todo
	}

	var node *registry.Node

	err = json.Unmarshal(body, &node)
	if err != nil {
		// todo
	}

	// , opts ...RegisterOption

	registry.Register(node)

	w.Write([]byte("OK"))
}
