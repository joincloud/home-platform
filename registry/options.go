package registry

import (
	"context"
	"time"
)

type Options struct {
}

type NodeOptions struct {
}

type RegisterOptions struct {
	TTL     time.Duration
	Context context.Context
}
