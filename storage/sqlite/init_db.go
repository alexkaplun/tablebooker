package sqlite

import (
	"math/rand"
	"strings"

	"github.com/alexkaplun/tablebooker/model"
	uuid "github.com/satori/go.uuid"
)

func (s *Storage) InitEmpty() error {
	// run the init queries
	for _, query := range strings.Split(sqlInitDatabase, ";") {
		if len(strings.TrimSpace(query)) == 0 {
			continue
		}

		_, err := s.db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) InitDummyData() error {

	// create 10 dummy tables without booking
	for i := 0; i < 10; i++ {
		table := model.Table{
			ID:     uuid.NewV4().String(),
			Number: i + 1,
			// capacity from 2 to 10
			Capacity: rand.Intn(8) + 2,
		}

		_, err := s.db.Exec(sqlInsertTable, table.ID, table.Number, table.Capacity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) InsertTable(table model.Table) error {
	_, err := s.db.Exec(sqlInsertTable, table.ID, table.Number, table.Capacity)
	if err != nil {
		return err
	}

	return nil
}
