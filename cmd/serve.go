/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	cmd "maestro/cmd/chaostests"
	"net/http"

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

		listenAddress, err := rootCmd.Flags().GetString("address")
		if err != nil {
			zap.L().Info(fmt.Sprintf("Error getting address:%s", err))
			return
		}
		zap.L().Info(fmt.Sprintf("Starting on %s", listenAddress))
		zap.L().Fatal(fmt.Sprint(http.ListenAndServe(listenAddress, r)))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

}

func RunTest(w http.ResponseWriter, r *http.Request) {

	var chaosTest ChaosTest
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var intermediate struct {
		Args []string `json:"args"`
	}

	err = json.Unmarshal(bodyBytes, &intermediate)
	if err != nil {
		log.Fatal(err)
	}

	// Check if there are enough elements in the args array
	if len(intermediate.Args) >= 2 {
		chaosTest = ChaosTest{
			Component: intermediate.Args[0],
			ChaosType: intermediate.Args[1],
		}

	}
	// Use the ChaosTest struct
	fmt.Printf("%+v\n", chaosTest)

	result := cmd.CheckList(chaosTest.Component, chaosTest.ChaosType)

	fmt.Fprintf(w, "\n Result: %s", result)

}

type ChaosTest struct {
	Component string
	ChaosType string
}
