package cmd

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/kingpin/v2"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	address     = kingpin.Flag("address", "The address and port for the server to listen to.").Envar("MAESTRO_LISTEN_ADDRESS").Default("0.0.0.0:8000").String()
	logLevel    = kingpin.Flag("log-level", "Log level (debug, info, warn, error, fatal, panic)").Envar("MAESTRO_LOG_LEVEL").Default("info").String()
	urlBasePath = kingpin.Flag("url-base-path", "The base URL to run Coroot at a sub-path, e.g. /base/").Envar("MAESTRO_URL_BASE_PATH").Default("/").String()
)

var version = "0.1.0"

type Options struct {
	BasePath string
	Version  string
}

func init() {
	kingpin.Version(version)
	kingpin.Parse()
	zap.ReplaceGlobals(createLogger(logLevel))
}

func main() {
	// Initialize Zap logger
	r := mux.NewRouter()
	r.PathPrefix("").Handler(http.StripPrefix(*urlBasePath, http.FileServer(http.Dir("./static"))))
	zap.L().Info(fmt.Sprintf("Starting on %s", *address))
	zap.L().Fatal(
		"Failed to start server",
		zap.Error(http.ListenAndServe(*address, r)),
	)
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
