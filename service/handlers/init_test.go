package handlers

import (
	"log"

	"github.com/alexkaplun/tablebooker/controller"
	"github.com/alexkaplun/tablebooker/storage/sqlite"
)

const DB_FILENAME = "test.db"

var c *controller.Controller
var s *sqlite.Storage

func init() {
	storage, err := sqlite.NewSQLiteStorage(DB_FILENAME)
	s = storage
	if err != nil {
		log.Fatal("error initializing storage", err)
	}

	storage.InitEmpty()

	c = controller.NewController(storage)
}
