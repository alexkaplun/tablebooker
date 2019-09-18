package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"

	"github.com/alexkaplun/tablebooker/service/common"

	"github.com/alexkaplun/tablebooker/model"
	"github.com/stretchr/testify/assert"
)

func TestBookTableByIdHandler_Negative(t *testing.T) {
	//create dummy table
	table := model.Table{
		ID:       uuid.NewV4().String(),
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
			id:             table.ID,
			book:           model.TableBook{BookDate: "", GuestName: "name", GuestContact: "contact"},
			expectedStatus: http.StatusBadRequest,
		},
		"empty guestName": {
			id:             table.ID,
			book:           model.TableBook{BookDate: today, GuestName: "", GuestContact: "contact"},
			expectedStatus: http.StatusBadRequest,
		},
		"empty guestContact": {
			id:             table.ID,
			book:           model.TableBook{BookDate: tomorrow, GuestName: "name", GuestContact: ""},
			expectedStatus: http.StatusBadRequest,
		},
		"bad date format": {
			id:             table.ID,
			book:           model.TableBook{BookDate: "2019-0101", GuestName: "name", GuestContact: "contact"},
			expectedStatus: http.StatusBadRequest,
		},
		"book date in the past": {
			id:             table.ID,
			book:           model.TableBook{BookDate: yesterday, GuestName: "name", GuestContact: "contact"},
			expectedStatus: http.StatusBadRequest,
		},
		"non-existing table": {
			id:             "Iamnonexistingtable",
			book:           model.TableBook{BookDate: tomorrow, GuestName: "name", GuestContact: "cotact"},
			expectedStatus: http.StatusNotFound,
		},
	}

	// prepare routes
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

// tests attempt to book over existing booking
// tests attempt to book on an empty date
func TestBookTableByIdHandler(t *testing.T) {
	//create dummy table
	table := model.Table{
		ID:       uuid.NewV4().String(),
		Number:   5,
		Capacity: 5,
	}
	//create a table
	err := s.InsertTable(table)
	assert.NoError(t, err)

	bookingDate := time.Now().UTC()

	bookingRequest := model.TableBook{
		TableID:      table.ID,
		BookDate:     bookingDate.Format("2006-01-02"),
		GuestName:    "Test Name",
		GuestContact: "Test Contact",
	}
	// book the table
	newBook, err := s.BookTableById(table.ID, bookingDate, bookingRequest.GuestName, bookingRequest.GuestContact)
	assert.NoError(t, err)
	if !assert.NotNil(t, newBook) {
		t.Fatal("got nil book from storage")
	}
	assert.NotEmpty(t, newBook.Code)

	// prepare routes
	r := chi.NewRouter()
	r.Route("/table", func(r chi.Router) {
		r.Post("/book/{table_id}", BookTableByIdHandler(c))
	})

	// attempt to book again on te same date
	reqBody, err := json.Marshal(bookingRequest)
	assert.NoError(t, err)

	makeReq := httptest.NewRequest(http.MethodPost, "/table/book/"+table.ID, bytes.NewReader(reqBody))
	out := httptest.NewRecorder()
	r.ServeHTTP(out, makeReq)

	// must be 409 status code
	assert.Equal(t, http.StatusConflict, out.Code)

	// now book on a next date
	bookingRequest.BookDate = bookingDate.Add(24 * time.Hour).Format("2006-01-02")
	reqBody, err = json.Marshal(bookingRequest)
	assert.NoError(t, err)

	makeReq = httptest.NewRequest(http.MethodPost, "/table/book/"+table.ID, bytes.NewReader(reqBody))
	out = httptest.NewRecorder()
	r.ServeHTTP(out, makeReq)

	assert.Equal(t, http.StatusOK, out.Code)

	// unmarshal response with TableBook object as a RawMessage
	var resp struct {
		Head *common.Head     `json:"head"`
		Body *json.RawMessage `json:"body"`
	}

	err = json.Unmarshal(out.Body.Bytes(), &resp)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	// unmarshal the bookResponse
	var bookResponse model.TableBook
	err = json.Unmarshal(*resp.Body, &bookResponse)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	// assert response values
	assert.Equal(t, bookingRequest.TableID, bookResponse.TableID)
	assert.NotEmpty(t, bookResponse.ID)
	assert.NotEmpty(t, bookResponse.Code)
	assert.Equal(t, bookingRequest.BookDate, bookResponse.BookDate)
	assert.Equal(t, bookingRequest.GuestName, bookResponse.GuestName)
	assert.Equal(t, bookingRequest.GuestContact, bookResponse.GuestContact)
}

func TestUnbookTableHandler(t *testing.T) {
	//create dummy table
	table := model.Table{
		ID:       uuid.NewV4().String(),
		Number:   5,
		Capacity: 7,
	}
	//create a table
	err := s.InsertTable(table)
	assert.NoError(t, err)

	// create custom booking
	bookingDate := time.Now().UTC()
	// book the table
	bookingRequest := model.TableBook{
		TableID:      table.ID,
		BookDate:     bookingDate.Format("2006-01-02"),
		GuestName:    "Test Name",
		GuestContact: "Test Contact",
	}
	newBook, err := s.BookTableById(table.ID, bookingDate, bookingRequest.GuestName, bookingRequest.GuestContact)
	assert.NoError(t, err)
	if !assert.NotNil(t, newBook) {
		t.Fatal("got nil book from storage")
	}
	assert.NotEmpty(t, newBook.Code)

	// prepare routes
	r := chi.NewRouter()
	r.Route("/table", func(r chi.Router) {
		r.Delete("/book/{code}", UnbookTableHandler(c))
	})

	// delete non-existing booking
	t.Run("non-existing booking", func(t *testing.T) {
		makeReq := httptest.NewRequest(http.MethodDelete, "/table/book/iamdummybookingcode", nil)
		out := httptest.NewRecorder()
		r.ServeHTTP(out, makeReq)
		assert.Equal(t, http.StatusNotFound, out.Code)
	})

	t.Run("valid booking code", func(t *testing.T) {
		makeReq := httptest.NewRequest(http.MethodDelete, "/table/book/"+newBook.Code, nil)
		out := httptest.NewRecorder()
		r.ServeHTTP(out, makeReq)
		// expect 200 OK
		assert.Equal(t, http.StatusOK, out.Code)

		out2 := httptest.NewRecorder()
		// now attempt to delete same code again
		r.ServeHTTP(out2, makeReq)
		// expect 404 Not found as booking must have been deleted
		assert.Equal(t, http.StatusNotFound, out2.Code)
	})
}
