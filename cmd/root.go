package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// vars to store provided filenames
var (
	image1 string
	image2 string
	output string
)

var RootCmd = &cobra.Command{
	Use:   "tablebooker <command>",
	Short: "Table Booker service",
	Long:  "Table Booker service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Table booker started")

	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
