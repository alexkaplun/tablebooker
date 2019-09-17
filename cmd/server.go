package cmd

import (
	"log"
	"net/http"

	"github.com/alexkaplun/tablebooker/service/handlers"

	"github.com/alexkaplun/tablebooker/controller"
	"github.com/alexkaplun/tablebooker/storage/sqlite"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start web server command",
	Long:  "Starts the server and listens for incoming requests",
	Run: func(cmd *cobra.Command, args []string) {

		storage, err := sqlite.NewSQLiteStorage(DB_FILENAME)
		if err != nil {
			log.Fatal("error initializing storage", err)
		}

		c := controller.NewController(storage)

		log.Println("Starting table booker web server on port 3000")
		log.Fatal(http.ListenAndServe(":3000", handlers.Routes(c)))
	},
}
