package controller

import (
	"log"
	"time"

	"github.com/alexkaplun/tablebooker/model"
)

func (c *Controller) GetTablesList() ([]*model.Table, error) {
	tables, err := c.storage.GetTablesList()
	if err != nil {
		log.Println("error getting tables from storage", err)
		return nil, err
	}
	return tables, nil
}

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

func (c *Controller) UnbookTableByCode(code string) (bool, error) {
	unbooked, err := c.storage.UnbookTableByCode(code)
	if err != nil {
		log.Println("error unbooking the table")
		return false, err
	}

	return unbooked, nil
}
