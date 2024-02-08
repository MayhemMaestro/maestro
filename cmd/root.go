/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/spf13/cobra"
)

var (
	// address = kingpin.Flag("address", "The address and port for the server to listen to.").Envar("MAESTRO_LISTEN_ADDRESS").Default("0.0.0.0:8000").String()

	logLevel = kingpin.Flag("log-level", "Log level (debug, info, warn, error, fatal, panic)").Envar("MAESTRO_LOG_LEVEL").Default("info").String()

	urlBasePath = kingpin.Flag("url-base-path", "The base URL to run Coroot at a sub-path, e.g. /base/").Envar("MAESTRO_URL_BASE_PATH").Default("/").String()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "maestro",
	Short: "A CLI tool to inject chaos into a Linux system",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("address", "", "The address for the server to listen on. Example: 0.0.0.0:8080")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
