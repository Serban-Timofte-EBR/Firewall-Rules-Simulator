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
	"fmt"
)

// SimulatePacket: Represents a single network packet being tested against the firewall rules.
type SimulatedPacket struct {
	SourceIP      string
	DestinationIP string
	Port          int
}

func main() {
	fmt.Println("Firewall Rules Simulator ... ")

	// Adding test rules
	rules.AddRule("192.168.1.1", "192.168.1.100", 80, rules.Allow)
	rules.AddRule("192.168.1.2", "192.168.1.101", 22, rules.Block)

	// Simulated packets list is mock data simulating real network packets.
	//packets := []SimulatedPacket{
	//	{"192.168.1.1", "192.168.1.100", 80},
	//	{"192.168.1.2", "192.168.1.101", 22},
	//	{"10.0.0.1", "10.0.0.2", 443},
	//}
	//
	//fmt.Println("Simulating packets ... ")
	//
	//for _, packet := range packets {
	//	action := rules.MatchRule(packet.SourceIP, packet.DestinationIP, packet.Port)
	//	if action == "" {
	//		fmt.Println("Packet from ", packet.SourceIP, " to ", packet.DestinationIP, " -> No rule matched!")
	//	} else {
	//		fmt.Println("Packet from ", packet.SourceIP, " to ", packet.DestinationIP, " - Action: ", action)
	//	}
	//}

	logFilePath := "firewall.log"
	logger.InitializeLogger(logFilePath)
	defer logger.CloseLogger()

	// Start capturing packets
	interfaceName := "en0"
	capture.StartCapture(interfaceName)
}
