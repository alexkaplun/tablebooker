package cmd

import (
	"log"

	"github.com/alexkaplun/tablebooker/controller"

	"github.com/alexkaplun/tablebooker/storage/sqlite"

	"github.com/spf13/cobra"
)

const DB_FILENAME = "./tables.db"

var initCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Init database command",
	Long:  "Inits the database with list of dummy tables",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting initdb")

		storage, err := sqlite.NewSQLiteStorage(DB_FILENAME)
		if err != nil {
			log.Println("error initializing storage", err)
			return
		}

		c := controller.NewController(storage)

		if err := c.InitDB(); err != nil {
			log.Println("error executing database init", err)
		} else {
			log.Println("initialize database successful")
		}
	},
}
