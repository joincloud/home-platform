package registry

import "github.com/joincloud/home-platform/registry/memory"

type Registry interface {
	Register(node *Node, opts ...RegisterOption) error
	Deregister(node *Node) error
	GetNode(...NodeOption) (node Node, err error)
	ListNodes(...NodeOption) (nodes []*Node, err error)
	Init(...Option) error
	Options() Options
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type NodeOption func(*NodeOptions)

var (
	reg = &memory.Registry{}
)

func Register(node *Node, opts ...RegisterOption) error {
	return reg.Register(node, opts...)
}
