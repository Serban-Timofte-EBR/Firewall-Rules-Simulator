# Development Roadmap

## Step 1: Initialize Project and Create File Structure
### What We Did
- Set up the directory structure for the Firewall Rules Simulator project.
- Created the following files and directories:
    - `main.go`: Entry point of the application.
    - `rules/rules.go`: Handles rule management.
    - `capture/capture.go`: Handles packet capture and analysis.
    - `logger/logger.go`: Manages logging.
    - `utils/utils.go`: Provides utility functions.

### Why We Did This
- To follow clean code principles by separating responsibilities into modules.
- To ensure the project is scalable and maintainable.

### Conex Info
- Each module will handle one core functionality:
    - `rules`: Rule definition and evaluation.
    - `capture`: Network packet capture and decoding.
    - `logger`: Logging for allowed and blocked traffic.
    - `utils`: Common utility functions shared across modules.

### Next Step
Start implementing the **rules module** to define and manage firewall rules.

## Step 2: Implement Basic Rule Management
### What We Did
- Implemented the foundation of the `rules` module.
- Created functions to:
  - Define rules (allow/block traffic based on IP and port).
  - Store rules in a slice for easy management.
  - Check if a given packet matches any rule.

### Why We Did This
- Rule management is a core functionality of the firewall simulator.
- This step allows us to begin filtering packets in later stages.

### Conex Info
- Rules are defined using structs to encapsulate properties like source IP, destination IP, and port.
- Functions in this step include:
  - `AddRule`: Add a new rule.
  - `MatchRule`: Check if a packet matches an existing rule.

### Next Step
Integrate rule-checking logic with a simulated packet to test functionality.

## Step 3: Simulate Packet Data for Testing Rules
### What We Did
- Added a mechanism to simulate packets with attributes like source IP, destination IP, and port.
- Integrated the packet simulation with the rules module to test whether packets are allowed or blocked.
- Printed the result of each simulated packet check to the console.

### Why We Did This
- Simulating packets allows us to verify that the rules system is working correctly without capturing real traffic.
- It’s a controlled way to debug and validate the filtering logic.

### Conex Info
- Simulated packets have the following attributes:
  - `SourceIP`
  - `DestinationIP`
  - `Port`
- Each packet is checked against the defined rules using the `MatchRule` function.
#### What is a Packet?
- A packet is a small unit of data transmitted over a network.
- When devices communicate over a network (e.g., sending an email, loading a webpage), the data is broken into smaller chunks called packets.
- Each packet contains:
  - **Header**: Metadata about the packet, including:
    - Source and destination IP addresses.
    - Port numbers.
    - Protocol information.
  - **Payload**: The actual data being transmitted (e.g., parts of a web page or file).

#### Role of Packets in a Firewall
- A firewall examines each packet passing through it to decide whether to **allow** or **block** it based on predefined rules.
- Key attributes analyzed in a packet include:
  - **Source IP Address**: The address of the device sending the packet.
  - **Destination IP Address**: The address of the intended recipient.
  - **Port**: The communication endpoint (e.g., port 80 for HTTP, port 22 for SSH).
  - **Protocol**: Specifies the communication type (e.g., TCP, UDP).
- Firewalls enforce security by:
  - Blocking unauthorized access to systems and sensitive data.
  - Allowing only specific types of traffic, such as web traffic on port 80 or secure connections on port 443.
- Packets that do not match any defined rule are typically handled using a **default action**, which can either be to allow or block them.

#### Importance
Understanding how packets work and their attributes is crucial for configuring firewall rules effectively and ensuring network security.
### Next Step
Extend the implementation to capture live packets using `gopacket`.

## Step 4: Capture Live Network Packets
### What We Did
- Integrated the `gopacket` library to capture live packets from the network interface.
- Decoded packet headers (source IP, destination IP, port) for evaluation.
- Passed captured packets to the rules module to check if they should be allowed or blocked.

### Why We Did This
- Capturing real packets allows the simulator to operate in a realistic environment.
- It bridges the gap between simulation and real-world firewall behavior.

### Conex Info
- **Packet Capture**:
  - Live packets are captured from a network interface using `gopacket` and `pcap`.
  - Requires permissions to access the network interface (e.g., root on Linux).
- **Packet Analysis**:
  - Captured packets are decoded to extract attributes like:
    - Source IP and destination IP.
    - Source and destination ports.
    - Protocol type (e.g., TCP/UDP).
- **Integration**:
  - The decoded packets are sent to the `MatchRule` function for rule evaluation.
- **Security**:
  - Packet capture tools must be used responsibly and only on authorized networks.

### Next Step
Implement logging for allowed and blocked packets and refine rule evaluation.