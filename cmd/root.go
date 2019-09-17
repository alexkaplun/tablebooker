package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tablebooker <command>",
	Short: "Table Booker service",
	Long:  "Table Booker service",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(serverCmd)
}
