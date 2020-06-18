package handler

import (
	"encoding/json"
	"github.com/joincloud/home-platform/util"
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

	util.RspJsonOK(w, "ok")
}

func ListNodes(w http.ResponseWriter, req *http.Request) {
	nodes, _ := registry.ListNodes()
	util.RspJsonOK(w, nodes)
}
