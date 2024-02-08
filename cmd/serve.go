/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var version = "0.1.0"

type Options struct {
	BasePath string
	Version  string
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Initialize the listening server",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		r := mux.NewRouter()
		r.PathPrefix("/static").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
		//
		r.HandleFunc("/chaos/tests/{chaosTest}", RunTest).Methods("POST")
		listenAddress := os.Getenv("MAESTRO_LISTEN_ADDRESS")
		zap.L().Info(fmt.Sprintf("Starting on %s", listenAddress))
		zap.L().Fatal(fmt.Sprint(http.ListenAndServe(listenAddress, r)))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

}

func RunTest(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	vars := mux.Vars(r)
	// Respond with the extracted parameters
	fmt.Fprintf(w, "You've requested the test: %s", vars["chaosTest"])
}
