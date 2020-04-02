package service

import "github.com/micro/go-micro/v2"

type Options struct {
	Service micro.Service
}

func MicroService(srv micro.Service) Option {
	return func(options *Options) {
		options.Service = srv
	}
}
