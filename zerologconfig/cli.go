package zerologconfig

import (
	"log/slog"
	"os"
	"slices"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	slogzerolog "github.com/samber/slog-zerolog/v2"
)

var (
	ValidLogLevels  = []string{"trace", "debug", "info", "warn", "error"}
	ValidLogFormats = []string{"plain", "color", "json"}
)

type LogConfig struct {
	LogLevel  string
	LogFormat string
	LogCaller bool
}

// Configure configures zerolog with the provided configuration
// Designed for use in CLI applications
func Configure(cfg LogConfig) {
	// log format
	if cfg.LogFormat == "" {
		cfg.LogFormat = "plain"
	} else if !slices.Contains(ValidLogFormats, cfg.LogFormat) {
		log.Error().Str("current", cfg.LogFormat).Strs("valid", ValidLogFormats).Msg("invalid log format specified, defaulting to plain")
		cfg.LogFormat = "plain"
	}
	var logContext zerolog.Context
	if cfg.LogFormat == "plain" {
		logContext = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true}).With().Timestamp()
	} else if cfg.LogFormat == "color" {
		colorableOutput := colorable.NewColorableStderr()
		logContext = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: colorableOutput, NoColor: false}).With().Timestamp()
	} else if cfg.LogFormat == "json" {
		logContext = zerolog.New(os.Stderr).Output(os.Stderr).With().Timestamp()
	}
	if cfg.LogCaller {
		logContext = logContext.Caller()
	}
	log.Logger = logContext.Logger()

	// log time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log level
	if !slices.Contains(ValidLogLevels, cfg.LogLevel) {
		log.Error().Str("current", cfg.LogLevel).Strs("valid", ValidLogLevels).Msg("invalid log level specified, defaulting to info")
		cfg.LogLevel = "info"
	}
	if cfg.LogLevel == "trace" {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else if cfg.LogLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if cfg.LogLevel == "info" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if cfg.LogLevel == "warn" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else if cfg.LogLevel == "error" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	// configure slog logger with zerolog adapter
	logger := slog.New(slogzerolog.Option{Level: zerologToSlogLevel(zerolog.GlobalLevel()), Logger: &log.Logger}.NewZerologHandler())
	slog.SetDefault(logger)

	// logging config
	log.Debug().Str("log-level", cfg.LogLevel).Str("log-format", cfg.LogFormat).Bool("log-caller", cfg.LogCaller).Msg("configured logging")
}

func zerologToSlogLevel(lvl zerolog.Level) slog.Level {
	switch lvl {
	case zerolog.TraceLevel:
		return slog.LevelDebug
	case zerolog.DebugLevel:
		return slog.LevelDebug
	case zerolog.InfoLevel:
		return slog.LevelInfo
	case zerolog.WarnLevel:
		return slog.LevelWarn
	case zerolog.ErrorLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
