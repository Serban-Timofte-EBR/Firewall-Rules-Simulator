// capture.go
// This module is responsible for capturing and analyzing network packets.
// Responsibilities include:
// - Intercepting packets from the network interface using gopacket.
// - Decoding packet headers (IP, TCP, UDP).
// - Passing packets to the rules module for filtering.

package capture
