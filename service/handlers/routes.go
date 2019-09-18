package handlers

import (
	"net/http"
	"time"

	"github.com/alexkaplun/tablebooker/controller"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(c *controller.Controller) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Route("/table", func(r chi.Router) {
		r.Post("/book/{table_id}", BookTableByIdHandler(c))
		r.Delete("/book/{code}", UnbookTableHandler(c))
		r.Get("/list", ListTableHandler(c))
	})

	return r
}
