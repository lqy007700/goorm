package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const (
	InfoLevel = iota
	ErrLevel
	Disabled
)

var (
	infoLog = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	errLog  = log.New(os.Stdout, "\033[31m[error ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers = []*log.Logger{errLog, infoLog}
	mu      sync.Mutex
)

var (
	Error  = errLog.Println
	Errorf = errLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrLevel < level {
		errLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
