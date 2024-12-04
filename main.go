// main.go
// This is the main entry point for the Firewall Rules Simulator application.
// It initializes the modules, sets up the environment, and coordinates the main functionalities:
// - Rule management
// - Packet capture
// - Traffic logging

package main

import (
	"Firewall-Rules-Simulator/capture"
	"Firewall-Rules-Simulator/logger"
	"Firewall-Rules-Simulator/rules"
	"flag"
	"fmt"
	"time"
)

// SimulatePacket: Represents a single network packet being tested against the firewall rules.
type SimulatedPacket struct {
	SourceIP      string
	DestinationIP string
	Port          int
}

func main() {
	fmt.Println("Firewall Rules Simulator ... ")

	// Define command-line arguments
	interfaceName := flag.String("interface", "en0", "The name of the network interface to capture packets from.")
	logFilePath := flag.String("logfile", "firewall.log", "The path to the log file.")
	defaultPolicy := flag.String("default-policy", "block", "The default policy for unmatched packets (allow/block).")
	duration := flag.Int("duration", 60, "The duration (in seconds) to capture packets.")
	configFilePath := flag.String("config", "config.json", "The path to the configuration file.")
	flag.Parse()

	// Initialize the logger
	logger.InitializeLogger(*logFilePath)
	defer logger.CloseLogger()

	// Set default policy
	switch *defaultPolicy {
	case "allow":
		rules.SetDefaultPolicy(rules.Allow)
	case "block":
		rules.SetDefaultPolicy(rules.Block)
	default:
		fmt.Println("Invalid default policy. Must be 'allow' or 'block'.")
		return
	}

	fmt.Printf("Starting Firewall Rules Simulator\n")
	fmt.Printf("Interface: %s\n", *interfaceName)
	fmt.Printf("Log File: %s\n", *logFilePath)
	fmt.Printf("Default Policy: %s\n", *defaultPolicy)

	// Reading from configuration file
	if err := rules.LoadConfig(*configFilePath); err != nil {
		fmt.Println("Error loading config file: ", err)
		return
	}

	// Start capturing packets
	if *duration > 0 {
		go capture.StartCapture(*interfaceName)
		time.Sleep(time.Duration(*duration) * time.Second)
		fmt.Println("Capture completed. Exiting ...")
	} else {
		capture.StartCapture(*interfaceName)
	}
}
