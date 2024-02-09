package cmd

import (
	"fmt"
	"os"

	conduct "github.com/MayhemMaestro/maestro/cmd/conduct"
	serve "github.com/MayhemMaestro/maestro/cmd/serve"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	cobra.OnInitialize(createLogger)
	rootCmd.PersistentFlags().String("log-level", "info", "The log level (debug, info, warn, error, fatal, panic)")
	rootCmd.PersistentFlags().String("encoding", "json", "The zap encoding to use (json or console)")

	viper.SetDefault("address", "0.0.0.0:8080")
	viper.SetDefault("log-level", "info")

	viper.SetConfigType("env")
	rootCmd.AddCommand(serve.ServeCmd)
	rootCmd.AddCommand(conduct.ConductCmd)

}

func createLogger() {
	var zapLevel zapcore.Level

	level, err := rootCmd.Flags().GetString("log-level")
	if err != nil {
		fmt.Println(err.Error())
		zap.L().Fatal("Failed to read listen log level", zap.Error(err))
	}
	encoding, err := rootCmd.Flags().GetString("encoding")
	if err != nil {
		fmt.Println(err.Error())
		zap.L().Fatal("Failed to read listen log level", zap.Error(err))
	}
	switch level {
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
		Encoding:          encoding,
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}

	logger := zap.Must(config.Build())
	zap.ReplaceGlobals(logger)
}
