# Firewall Rules Simulator

## Project Overview
The **Firewall Rules Simulator** is a GoLang-based application designed to simulate the behavior of a basic firewall. It allows users to define and apply filtering rules for network packets based on IP addresses, ports, and protocols (TCP/UDP). The application will capture network traffic, analyze it against defined rules, and log whether packets are allowed or blocked.

---

## Objectives
1. **Packet Filtering**: Simulate a firewall by filtering network packets based on user-defined rules.
2. **Packet Capture and Analysis**: Use GoLang libraries to capture live traffic and decode packet headers.
3. **Logging**: Maintain logs of all filtered traffic for auditing and debugging purposes.
4. **Learning**: Gain hands-on experience with networking, GoLang programming, and cybersecurity fundamentals.

---

## Key Features
- Define rules based on:
    - Source and destination IP addresses.
    - Ports (e.g., 80 for HTTP, 443 for HTTPS).
    - Protocols (TCP/UDP).
- Capture live network traffic for analysis.
- Filter packets in real-time and log the results.
- Extendable architecture for future enhancements like stateful filtering.

---

## Technologies Used
- **Programming Language**: GoLang
- **Libraries**:
    - [gopacket](https://github.com/google/gopacket): For packet capture and decoding.
    - **Standard Go Libraries**:
        - `net`, `net/http`: For networking functionality.
        - `log`: For logging blocked and allowed packets.
- **Concurrency**: Go’s goroutines and channels for efficient handling of multiple packet streams.

---

## Learning Goals
- **Networking Concepts**:
    - IP, TCP, UDP, and packet filtering.
    - How firewalls process network traffic.
- **Packet Analysis**:
    - Learn to decode and filter packets using the gopacket library.

---

## Development Plan
1. **Set Up Environment**:
    - Install GoLang and required libraries (gopacket, etc.).
    - Configure the development tools and dependencies.
2. **Define Application Workflow**:
    - User interface for defining rules.
    - Logic for filtering packets and logging.
3. **Incremental Development**:
    - Start with basic packet capture.
    - Add filtering logic.
    - Implement logging and user-defined rules.
4. **Testing and Documentation**:
    - Test on a controlled network.
    - Document learning outcomes and implementation steps.

---

## Getting Started
### Prerequisites
- **GoLang**: [Install GoLang](https://golang.org/doc/install)
- **gopacket Library**: Install using `go get github.com/google/gopacket`