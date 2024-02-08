/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
		// Extract path parameters

		listenAddr, _ := cmd.Flags().GetString("address")
		zap.L().Info(fmt.Sprintf("Starting on %s", listenAddr))
		zap.L().Fatal(fmt.Sprint(http.ListenAndServe(listenAddr, r)))
	},
}

func init() {
	zap.ReplaceGlobals(createLogger(logLevel))
	rootCmd.AddCommand(serveCmd)

}

func RunTest(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	vars := mux.Vars(r)
	test := vars["chaosTest"]
	// Respond with the extracted parameters
	fmt.Fprintf(w, "You've requested the test: %s", test)
}

func createLogger(level *string) *zap.Logger {
	var zapLevel zapcore.Level
	switch *level {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "fatal":
		zapLevel = zap.FatalLevel
	case "panic":
		zapLevel = zap.PanicLevel
	default:
		zapLevel = zap.InfoLevel
	}
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}

	return zap.Must(config.Build())
}
