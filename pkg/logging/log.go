package logging

import (
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"

	"PennyHardway/pkg/file"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	//logger *log.Logger
	logger     zerolog.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	//logger = log.New(F, DefaultPrefix, log.LstdFlags)
	output := zerolog.ConsoleWriter{Out: F, TimeFormat: time.RFC850}
	logger = zerolog.New(output).With().Caller().Timestamp().Logger()

	//logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	logger.Debug().Msgf("%v", v)
}

func Info(v ...interface{}) {
	logger.Info().Msgf("%v", v)
}

func Warn(v ...interface{}) {
	logger.Warn().Msgf("%v", v)
}

func Error(v ...interface{}) {
	logger.Error().Msgf("%v", v)
}

func Fatal(v ...interface{}) {
	logger.Fatal().Msgf("%v", v)
}
