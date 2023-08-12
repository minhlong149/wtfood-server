package util

import "log"

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type LogService struct {
	Logger *log.Logger
}

func SetupLogger() *LogService {
	return &LogService{Logger: log.Default()}
}

func (log *LogService) Info(args ...interface{}) {
	log.Logger.Printf("%v", args...)
}

func (log *LogService) Error(args ...interface{}) {
	log.Logger.Fatalf("%v", args...)
}
