package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexkaplun/tablebooker/model"
	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi"
)

func TestBookTableByIdHandler_Negative(t *testing.T) {
	//create dummy table
	id := "2d06220d-962f-4ab0-8847-78e9e3b45d83"
	table := model.Table{
		ID:       id,
		Number:   5,
		Capacity: 5,
	}
	//create a table
	s.InsertTable(table)

	today := time.Now().UTC().Format("2006-01-02")
	tomorrow := time.Now().UTC().Add(24 * time.Hour).Format("2006-01-02")
	yesterday := time.Now().UTC().Add(-24 * time.Hour).Format("2006-01-02")

	tests := map[string]struct {
		id             string
		book           model.TableBook
		expectedStatus int
	}{
		"empty bookDate": {
			id:             id,
			book:           model.TableBook{BookDate: "", GuestName: "name", GuestContact: "contact"},
			expectedStatus: http.StatusBadRequest,
		},
		"empty guestName": {
			id:             id,
			book:           model.TableBook{BookDate: today, GuestName: "", GuestContact: "contact"},
			expectedStatus: http.StatusBadRequest,
		},
		"empty guestContact": {
			id:             id,
			book:           model.TableBook{BookDate: tomorrow, GuestName: "name", GuestContact: ""},
			expectedStatus: http.StatusBadRequest,
		},
		"bad date format": {
			id:             id,
			book:           model.TableBook{BookDate: "2019-0101", GuestName: "name", GuestContact: "contact"},
			expectedStatus: http.StatusBadRequest,
		},
		"book date in the past": {
			id:             id,
			book:           model.TableBook{BookDate: yesterday, GuestName: "name", GuestContact: "contact"},
			expectedStatus: http.StatusBadRequest,
		},
		"non-existing table": {
			id:             "Iamnonexistingtable",
			book:           model.TableBook{BookDate: tomorrow, GuestName: "name", GuestContact: "cotact"},
			expectedStatus: http.StatusNotFound,
		},
	}

	r := chi.NewRouter()
	r.Route("/table", func(r chi.Router) {
		r.Post("/book/{table_id}", BookTableByIdHandler(c))
	})

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reqBody, err := json.Marshal(test.book)
			assert.NoError(t, err)

			makeReq := httptest.NewRequest(http.MethodPost, "/table/book/"+test.id, bytes.NewReader(reqBody))
			out := httptest.NewRecorder()
			r.ServeHTTP(out, makeReq)

			assert.Equal(t, test.expectedStatus, out.Code)
		})
	}

}
