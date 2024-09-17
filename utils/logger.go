package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "secretsafe: ", log.LstdFlags)

// Info logs an informational message
func Info(msg string) {
	logger.Println("[INFO]", msg)
}

// Error logs an error message
func Error(msg string) {
	logger.Println("[ERROR]", msg)
}
