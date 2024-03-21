package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MayhemMaestro/maestro/chaostests"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ChaosTest struct {
	Component      string
	ChaosType      string
	ChaosThreshold string
	ChaosLength    string
}

// ServeCmd represents the serve command
var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "run the http server",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		addr, exists := os.LookupEnv("MAESTRO_LISTEN_ADDRESS")
		if !exists {
			addr = "localhost:8080"
		}

		r := mux.NewRouter()
		//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

		r.HandleFunc("/chaos/tests/{chaosTest}", RunTest).Methods("POST")

		// listenAddress, err := ServeCmd.Flags().GetString("address")
		// if err != nil {
		// 	zap.L().Fatal("Failed to read listen address", zap.Error(err))
		// }
		server := &http.Server{
			Addr:              addr,
			ReadHeaderTimeout: 10 * time.Second,
			Handler:           r,
		}
		zap.L().Info(fmt.Sprintf("Starting on %s", addr))
		zap.L().Fatal(
			"Failed to start server",
			zap.Error(server.ListenAndServe()),
		)
	},
}

func init() {
	viper.SetDefault("serveAddress", "localhost:8080")

	ServeCmd.Flags().String("address", viper.GetString("serveAddress"), "The address for the server to listen on. Example: 0.0.0.0:8080")
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

	if len(intermediate.Args) >= 2 {
		chaosTest = ChaosTest{
			Component:      intermediate.Args[0],
			ChaosType:      intermediate.Args[1],
			ChaosThreshold: intermediate.Args[2],
			ChaosLength:    intermediate.Args[3],
		}

	}
	fmt.Printf("%+v\n", chaosTest)

	result := chaostests.CheckList(chaosTest.Component, chaosTest.ChaosType, chaosTest.ChaosThreshold, chaosTest.ChaosLength)

	fmt.Fprintf(w, "\n Result: %s", result)

}
