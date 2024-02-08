/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	//address = kingpin.Flag("address", "The address and port for the server to listen to.").Envar("MAESTRO_LISTEN_ADDRESS").Default("0.0.0.0:8000").String()

	logLevel = kingpin.Flag("log-level", "Log level (debug, info, warn, error, fatal, panic)").Envar("MAESTRO_LOG_LEVEL").Default("info").String()

	// urlBasePath = kingpin.Flag("url-base-path", "The base URL to run Coroot at a sub-path, e.g. /base/").Envar("MAESTRO_URL_BASE_PATH").Default("/").String()
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
	rootCmd.PersistentFlags().String("address", "0.0.0.0:8080", "The address for the server to listen on. Example: 0.0.0.0:8080")
	zap.ReplaceGlobals(createLogger(logLevel))

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.SetDefault("address", "0.0.0.0:8080")

	viper.SetConfigType("env")
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
