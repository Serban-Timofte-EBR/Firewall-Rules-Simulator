// capture.go
// This module is responsible for capturing and analyzing network packets.
// Responsibilities include:
// - Intercepting packets from the network interface using gopacket.
// - Decoding packet headers (IP, TCP, UDP).
// - Passing packets to the rules module for filtering.

package capture

import (
	"Firewall-Rules-Simulator/rules"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
)

func StartCapture(interfaceName string) {
	// Step 1: Open the device for packet capture
	// device (interfaceName): The name of the device to open / Of the network interface (e.g., eth0 - Ethernet, wlan0 - Wi-Fi).
	// 1600 the maximum number of bytes to read from each packet.
	// true: promiscuous mode (capture all packets on the network).
	// pcap.BlockForever: timeout (wait indefinitely for packets).
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Error opening device: %v", err)
	}
	defer handle.Close()

	fmt.Println("Capturing packets for " + interfaceName + " ...")

	// Step 2: Create a PacketSource to read packets from the handle.
	// - `gopacket.NewPacketSource` wraps the `handle` and parses the packets into higher-level structures.
	packageSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Step 3: Process each captured packet.
	// - The `Packets()` channel streams captured packets one at a time.
	// - Each packet is passed to the `processPacket` function for further analysis and rule checking.
	for packet := range packageSource.Packets() {
		processPacket(packet)
	}
}

// processPacket analyzes and extracts information from a captured packet.
// Parameters:
//
//	packet (gopacket.Packet): The packet captured by the network interface.
//
// Description:
//
//	This function processes a packet by:
//	1. Extracting the network layer (e.g., IP addresses).
//	2. Extracting the transport layer (e.g., TCP/UDP ports).
//	3. Passing extracted attributes (source IP, destination IP, and port) to the `checkPacket` function.
//
// Steps:
//  1. Get the network layer of the packet to extract source and destination IPs.
//  2. Get the transport layer to identify the protocol (TCP/UDP) and extract port numbers.
//  3. Pass the extracted details to `checkPacket` for rule evaluation.
func processPacket(packet gopacket.Packet) {
	// Step 1: Extract the network layer (IP addresses).
	networkLayer := packet.NetworkLayer()
	if networkLayer == nil {
		return
	}
	srcIP := networkLayer.NetworkFlow().Src().String()
	destIP := networkLayer.NetworkFlow().Dst().String()

	// Step 2: Extract the transport layer (TCP/UDP ports).
	transportLayer := packet.TransportLayer()
	if transportLayer == nil {
		return
	}

	// Identify the transport protocol (TCP/UDP) and extract port numbers.
	switch packet.TransportLayer().LayerType() {
	case layers.LayerTypeTCP:
		tcp, ok := transportLayer.(*layers.TCP)
		if tcp == nil || !ok {
			return
		}
		checkPacket(srcIP, destIP, int(tcp.DstPort))

	case layers.LayerTypeUDP:
		if udp, ok := transportLayer.(*layers.UDP); ok {
			checkPacket(srcIP, destIP, int(udp.DstPort))
		}
	}
}

// checkPacket evaluates a packet's details (source IP, destination IP, and port)
// against the firewall rules and determines the action to take.
// Parameters:
//
//	srcIP (string): The source IP address of the packet.
//	dstIP (string): The destination IP address of the packet.
//	port (int): The destination port of the packet.
//
// Description:
//   - Calls the `MatchRule` function from the `rules` package to evaluate the packet.
//   - Logs the action ("allow", "block", or "no match") based on the matching rule.
//   - This function serves as the decision-maker for how to handle the packet.
func checkPacket(srcIP string, destIP string, i int) {
	// Step 1: Evaluate the packet against the firewall rules.
	action := rules.MatchRule(srcIP, destIP, i)

	// Step 2: Log the action based on the rule match.
	if action == "" {
		fmt.Println("Packet from ", srcIP, " to ", destIP, " -> No rule matched!")
	} else {
		fmt.Println("Packet from ", srcIP, " to ", destIP, " - Action: ", action)
	}
}
