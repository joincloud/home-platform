package service

import "github.com/micro/go-micro/v2"

type Service interface {
	Options() Options
}

type Option func(*Options)

type DefaultService struct {
	options Options
	Service micro.Service
}

func (s *DefaultService) Options() Options {
	return s.options
}

func NewService(options ...Option) Service {
	opts := Options{}

	for _, o := range options {
		o(&opts)
	}

	if opts.Service == nil {
		opts.Service = micro.NewService()
	}

	return &DefaultService{
		options: opts,
		Service: opts.Service,
	}
}
