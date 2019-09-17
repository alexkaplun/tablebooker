package controller

import "github.com/alexkaplun/tablebooker/storage/sqlite"

type Controller struct {
	storage *sqlite.Storage
}

func NewController(storage *sqlite.Storage) *Controller {
	c := new(Controller)
	c.storage = storage
	return c
}
