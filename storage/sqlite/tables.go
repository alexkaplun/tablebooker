package sqlite

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/alexkaplun/tablebooker/model"
)

func (s *Storage) GetAvailableTables() ([]*model.Table, error) {
	rows, err := s.db.Query(sqlSelectAllTables)
	if err != nil {
		return nil, err
	}
	tables := make([]*model.Table, 0)
	for rows.Next() {
		item := model.Table{}
		if err := rows.Scan(&item.ID, &item.Number, &item.Capacity); err != nil {
			return nil, err
		}
		tables = append(tables, &item)
	}
	return tables, nil
}

func (s *Storage) CheckTableExists(id string) (bool, error) {
	rows, err := s.db.Query(sqlTableExists, id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, nil
	}
	return true, nil
}

func (s *Storage) IsTableAvailable(id string, bookDate time.Time) (bool, error) {
	rows, err := s.db.Query(sqlTableBookedOnDate, id, bookDate)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// booking found on date
	if rows.Next() {
		return false, nil
	}

	return true, nil
}

func (s *Storage) BookTableById(tableId string, bookDate time.Time, guestName string, guestContact string) (*model.TableBook, error) {
	// generate the book id
	id := uuid.NewV4().String()
	// generate the code
	code := uuid.NewV4().String()

	_, err := s.db.Exec(sqlCreateBook, id, tableId, bookDate, guestName, guestContact, code)
	if err != nil {
		return nil, err
	}

	return &model.TableBook{
		ID:           id,
		TableID:      tableId,
		BookDate:     bookDate.Format("2006-01-02"),
		GuestName:    guestName,
		GuestContact: guestContact,
		Code:         code,
	}, nil
}
