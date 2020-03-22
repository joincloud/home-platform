package registry

type Registry interface {
	Register(node *Node, opts ...RegisterOption)
	Deregister(node *Node)
	GetNodes(...NodeOption)
	Init(...Option) error
	Options() Options
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type NodeOption func(*NodeOptions)
