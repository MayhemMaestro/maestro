package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// ConductCmd represents the conduct command
var ConductCmd = &cobra.Command{
	Use:   "conduct [list]",
	Short: "Initiates a new chaos test",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {

			fmt.Println("Error: Please provide the name of the component to test [e.g. 'cpu' or 'mem']")
			return
		}

		if len(args) < 2 {

			fmt.Println("Error: Please provide the name of the test type. run 'list <component>' to view all tests for a component.")
			return
		}

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

		if err != nil {
			zap.L().Info(fmt.Sprintf("Error getting address:%s", err))
			return
		}

		url := "http://" + "listenAddress" + "/chaos/tests/" + args[0]

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

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			zap.L().Error("Error reading response", zap.Error(err))
		}
		zap.L().Info(fmt.Sprintf("Response: %s", body))
	},
}

var conductSubCmd = &cobra.Command{
	Use:   "list [cpu, mem]",
	Short: "My subcommand",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {

			fmt.Println("All components: [cpu, mem, net]. Specify 'list  and then a component to view relevant tests")
			return
		}

		if args[0] == "cpu" {
			fmt.Printf("[saturation, latency]")
		}
		if args[0] == "mem" {
			fmt.Printf("[oom, corruption]")
		}
	},
}

func init() {
	ConductCmd.AddCommand(conductSubCmd)

	ConductCmd.SetArgs([]string{"list"})

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ConductCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ConductCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
