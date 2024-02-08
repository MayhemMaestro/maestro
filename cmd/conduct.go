/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// conductCmd represents the conduct command
var conductCmd = &cobra.Command{
	Use:   "conduct",
	Short: "Initiates a new chaos test",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		data := map[string]interface{}{
			"message": "maestro conduct called",
			"args":    args,
		}

		// Convert the data to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling data:", err)
			return
		}

		// Create a new HTTP client
		client := &http.Client{}
		if err != nil {
			zap.L().Info(fmt.Sprintf("Error getting address:%s", err))
			return
		}
		listenAddress := os.Getenv("MAESTRO_LISTEN_ADDRESS")

		var url string

		switch args[0] {
		case "cpu":
			url = "http://" + listenAddress + "/chaos/tests/cpu"

		case "mem":
			url = "http://" + listenAddress + "/chaos/tests/mem"

		default:
			url = "http://" + listenAddress + "/chaos/tests/help"
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			zap.L().Info(fmt.Sprintf("Error creating request:%s", err))
			return
		}

		// Set the Content-Type header to application/json
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			zap.L().Info(fmt.Sprintf("Error sending request: %s", err))
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		zap.L().Info(fmt.Sprintf("Response: %s", body))
		// Print the response status code
	},
}

func init() {

	rootCmd.AddCommand(conductCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// conductCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// conductCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
