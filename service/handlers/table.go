package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexkaplun/tablebooker/model"

	"github.com/go-chi/chi"

	"github.com/alexkaplun/tablebooker/service/common"

	"github.com/alexkaplun/tablebooker/controller"
)

func BookTableByIdHandler(c *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tableId := chi.URLParam(r, "table_id")
		// validate tableId
		if tableId == "" {
			http.Error(w, "empty table id", http.StatusBadRequest)
			return
		}

		tableBook := model.TableBook{}
		// decode the request
		err := json.NewDecoder(r.Body).Decode(&tableBook)
		if err != nil {
			http.Error(w, "error parsing request", http.StatusBadRequest)
			return
		}

		// validate incoming parameters
		if tableBook.BookDate == "" {
			http.Error(w, "empty book date", http.StatusBadRequest)
			return
		}

		if tableBook.GuestName == "" {
			http.Error(w, "empty guest name", http.StatusBadRequest)
			return
		}

		if tableBook.GuestContact == "" {
			http.Error(w, "empty guest contact", http.StatusBadRequest)
			return
		}

		// validate the bookdate format
		bookDate, err := time.Parse("2006-01-02", tableBook.BookDate)
		if err != nil {
			http.Error(w, "bad bookDate format", http.StatusBadRequest)
			return
		}

		// check book date is today or in future
		// consider times are all UTC
		now := time.Now().UTC()
		if time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).
			After(bookDate) {
			http.Error(w, "bookDate is in the past", http.StatusBadRequest)
			return
		}

		// attempt to book the table
		bookResult, err := c.BookTableById(tableId, tableBook)
		if err != nil {
			http.Error(w, "error booking table", http.StatusInternalServerError)
			return
		}

		// book result does not exist
		if bookResult == nil {
			http.Error(w, "table not found", http.StatusNotFound)
			return
		}

		// bookResult.code empty, table is already booked
		if bookResult.Code == "" {
			http.Error(w, "table already booked", http.StatusConflict)
			return
		}

		common.NewResponse(bookResult).WriteResponse(w)
	}
}

func UnbookTableHandler(c *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func ListTableHandler(c *controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tables, err := c.GetAvailableTables()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		common.NewResponse(tables).WriteResponse(w)
	}
}
