package main

import (
	"log"
	"os"
)

func Logger() *log.Logger {
	const logFilePath = "apiLog.log"
	file, _ := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger := log.New(file, "", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	return logger
}
