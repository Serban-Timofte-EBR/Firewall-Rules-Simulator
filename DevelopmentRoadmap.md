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

## Step 5: Add Logging for Packet Activity
### What We Did
- Integrated logging to record the details of processed packets.
- Logged information includes:
  - Timestamp
  - Source IP
  - Destination IP
  - Port
  - Action (allow/block/no match)

### Why We Did This
- Logging provides a persistent record for auditing and debugging.
- Helps identify potential misconfigurations in firewall rules.

### Conex Info
- **Logging Levels**:
  - Logs can be written to a file or displayed in the terminal.
  - Different logging levels (e.g., info, error) can help categorize messages.
- **Structure**:
  - Each log entry includes:
    - Timestamp for when the packet was processed.
    - Source and destination IPs, port, and evaluation result.
- **Future Enhancements**:
  - Integrate with external log monitoring tools.
  - Add verbosity levels to control the amount of logged information.

## Step 6: Enhance Rule Management
### What We Did
- Added support for:
  - Protocol-specific rules (e.g., allow/block only TCP or UDP traffic).
  - Wildcard rules (e.g., block all traffic to a specific port, regardless of IP).
  - Dynamic rule updates (add/remove rules without restarting the program).

### Why We Did This
- Protocol-specific rules make the firewall more granular and realistic.
- Wildcard rules simplify management by covering broader scenarios.
- Dynamic updates improve usability by allowing changes during runtime.

### Conex Info
- **Protocol-Specific Rules**:
  - Include `TCP`, `UDP`, and `ANY`

## Step 7: Add Default Policy
### What We Did
- Implemented a default policy for unmatched packets.
- Added a mechanism to configure the default policy (allow or block).
- Integrated the default policy into the packet evaluation logic.

### Why We Did This
- Ensures that the firewall handles all packets, even those that do not match any rule.
- Mimics real-world firewall behavior, where a default action is applied to unmatched traffic.

### Conex Info
- **Default Policy**:
  - Can be configured to "allow" or "block".
  - Acts as a fallback when no rule matches.
- **Example Use Cases**:
  - Default "block" for strict security.
  - Default "allow" for a permissive setup.

## Step 8: Add Support for Command-Line Arguments (CLI)
### What We Did
- Integrated a command-line interface to allow runtime configuration of:
  - Network interface.
  - Log file path.
  - Default policy (allow/block).
  - Optional capture duration.
- Used the `flag` package for argument parsing.

### Why We Did This
- A CLI improves usability and flexibility, allowing users to:
  - Specify settings dynamically at runtime.
  - Avoid hardcoding values into the source code.

### Conex Info
- **CLI Basics**:
  - The `flag` package is part of Go's standard library for parsing command-line arguments.
  - Each argument is defined with a name, default value, and description.
- **Examples**:
  - Run with specific options:
    ```bash
    sudo go run main.go --interface en0 --logfile firewall.log --default-policy block --duration 10
    ```
  - Default policy: Configures unmatched packets to be allowed or blocked.
  - Capture duration: Specifies how long the program runs (e.g., 10 seconds).

## Step 9: Integrate Configuration File Support
### What We Did
- Implemented the ability to load firewall rules and settings from a configuration file (`config.json`).
- Added support for:
  - Default policy configuration.
  - Rule definitions (source/destination IP, port, protocol, action).

### Why We Did This
- Simplifies the process of defining and managing firewall rules.
- Allows changes to rules and settings without modifying the code or recompiling.

### Conex Info
- **Configuration File Format**:
  - The configuration file is in JSON format for readability and ease of parsing.
  - Example structure:
    ```json
    {
      "defaultPolicy": "block",
      "rules": [
        { "sourceIP": "192.168.1.1", "destinationIP": "*", "port": 80, "protocol": "TCP", "action": "allow" },
        { "sourceIP": "*", "destinationIP": "192.168.1.101", "port": 22, "protocol": "TCP", "action": "block" }
      ]
    }
    ```
- **Dynamic Reload**:
  - Future enhancements can include live reloading of configurations at runtime.

## Step 10: Add Unit Tests for Core Functions
### What We Did
- Implemented unit tests for:
  - Rule matching logic.
  - Loading configuration from a JSON file.
  - Logging traffic to the file.
- Used Go's `testing` package to write and execute the tests.

### Why We Did This
- To ensure core functionalities behave as expected.
- To catch potential bugs during development.

### Conex Info
- **Testing Framework**:
  - Go’s standard `testing` package is used for writing and running tests.
- **Mock Data**:
  - Test cases use mock data to simulate various scenarios (e.g., matching rules, loading valid and invalid configs).