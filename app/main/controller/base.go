package controller

import (
	cake "github.com/pobyzaarif/cake_store/business/cake"
)

type Controller struct {
	cakeService cake.Service
}

func NewController(cakeService cake.Service) *Controller {
	return &Controller{
		cakeService,
	}
}
