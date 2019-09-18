package controller

import (
	"log"
	"time"

	"github.com/alexkaplun/tablebooker/model"
)

// return list of tables in database
func (c *Controller) GetTablesList() ([]*model.Table, error) {
	tables, err := c.storage.GetTablesList()
	if err != nil {
		log.Println("error getting tables from storage", err)
		return nil, err
	}
	return tables, nil
}

// attempts to book the table by its id
// will return pointer to model.TableBook object,containing all info about booking
func (c *Controller) BookTableById(id string, book model.TableBook) (*model.TableBook, error) {
	bookDate, err := time.Parse("2006-01-02", book.BookDate)
	if err != nil {
		log.Println("error parsing the book date", err)
		return nil, err
	}

	// check if table exists
	exists, err := c.storage.CheckTableExists(id)
	if err != nil {
		log.Println("error checking table existence", err)
		return nil, err
	}

	// return nil if the table does not exist
	if !exists {
		return nil, nil
	}

	// check if table is available on date
	isAvailable, err := c.storage.IsTableAvailable(id, bookDate)
	if err != nil {
		log.Println("error checking table availability", err)
		return nil, err
	}

	// return empty table book if table is not available
	if !isAvailable {
		return &model.TableBook{}, nil
	}

	bookResult, err := c.storage.BookTableById(id, bookDate, book.GuestName, book.GuestContact)
	if err != nil {
		log.Println("error creating new book", err)
		return nil, err
	}

	return bookResult, nil
}

// attempts to unbook the table
// returns false if booking was not found
// returns true if onbook was successful
func (c *Controller) UnbookTableByCode(code string) (bool, error) {
	unbooked, err := c.storage.UnbookTableByCode(code)
	if err != nil {
		log.Println("error unbooking the table")
		return false, err
	}

	return unbooked, nil
}
