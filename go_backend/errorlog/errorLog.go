package errorlog

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	errorLog *log.Logger
}

var instance *logger
var once sync.Once

func GetLogger() *logger {
	once.Do(func() {
		// Open error log file
		file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open error log file: %s", err)
		}

		// Create new logger instance
		instance = &logger{
			errorLog: log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})

	return instance
}

func (l *logger) LogError(err error) {
	l.errorLog.Println(err)
}
