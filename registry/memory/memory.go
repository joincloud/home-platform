package memory

import (
	"fmt"
	"sync"

	"github.com/joincloud/home-platform/registry"
)

type Registry struct {
	options registry.Options

	sync.RWMutex
	nodes map[string]*registry.Node
}

func (r *Registry) Register(node *registry.Node, opts ...registry.RegisterOption) error {
	r.Lock()
	defer r.Unlock()

	var options registry.RegisterOptions
	for _, o := range opts {
		o(&options)
	}

	if _, ok := r.nodes[node.Name]; ok {
		return fmt.Errorf("registry the node %s existed", node.Name)
	}

	r.nodes[node.Name] = node

	return nil
}

func (r *Registry) Deregister(node *registry.Node) {
	panic("implement me")
}

func (r *Registry) GetNodes(option ...registry.NodeOption) {
	panic("implement me")
}

func (r *Registry) Options() registry.Options {
	panic("implement me")
}

func (r *Registry) Init(opts ...registry.Option) error {
	for _, o := range opts {
		o(&r.options)
	}

	r.Lock()
	defer r.Unlock()

	// todo load nodes from db or file

	return nil
}
