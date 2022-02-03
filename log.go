package icap

import (
	"io/ioutil"
	"log"
	"os"
)

type loggerType int

const (
	debugg loggerType = iota
	standard
)

func newLogger(lType loggerType) *log.Logger {
	if lType == debugg {
		if debugEnvIsSet := os.Getenv("DEBUG"); debugEnvIsSet != "" {
			logger := log.New(os.Stdout, "DEBUG ", log.LstdFlags)
			logger.Println("DEBUGGING ENABLED")
			return logger
		}
		return log.New(ioutil.Discard, "DEBUG ", log.LstdFlags)
	}
	return log.New(os.Stderr, "", log.LstdFlags)
}

var (
	Debug = newLogger(debugg)
	Std   = newLogger(standard)
)
