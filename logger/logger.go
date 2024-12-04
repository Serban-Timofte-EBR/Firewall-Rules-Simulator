// logger.go
// This module handles logging for the Firewall Rules Simulator.
// Responsibilities include:
// - Logging allowed and blocked traffic.
// - Writing logs to a file or displaying them on the console.
// - Providing a mechanism for debugging and auditing.
package logger

import (
	"log"
	"os"
	"time"
)

var logFile *os.File

func InitializeLogger(filePath string) {
	var err error
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	log.SetOutput(logFile)
	log.Printf("Logger initialized. Logging to file: %s\n", filePath)
}

func LogTraffic(srcIP, destIP string, port int, action string) {
	timestamp := time.Now().Format(time.RFC850)
	log.Printf("[%s] Action: %s | Source: %s | Destination: %s | Port: %d\n", timestamp, action, srcIP, destIP, port)
}

func CloseLogger() {
	if logFile != nil {
		log.Printf("Closing log file.")
		logFile.Close()
	}
}
