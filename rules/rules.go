// rules.go
// This module handles the management of firewall rules.
// Responsibilities include:
// - Defining rules (e.g., allow/block based on IP, port, protocol).
// - Storing and updating rule sets.
// - Checking packets against defined rules.

package rules

import "fmt"

type Action string

const (
	Allow Action = "allow"
	Block Action = "block"
)

// Firewall rule
type Rule struct {
	SourceIP      string
	DestinationIP string
	Port          int
	Protocol      string
	Action        Action
}

var rules []Rule
var defaultPolicy Action = Block // Default policy is to block unmatched packets

// SetDefaultPolicy sets the default action for unmatched packets
func SetDefaultPolicy(action Action) {
	if action != Allow && action != Block {
		fmt.Println("Invalid action: %s. Must be 'allow', or 'block'.", action)
		return
	}
	defaultPolicy = action
	fmt.Println("Default policy set to: ", action)
}

// GetDefaultPolicy returns the default action for unmatched packets
func GetDefaultPolicy() Action {
	return defaultPolicy
}

// AddRule adds a new rule to the list
func AddRule(sourceIP, destinationIP string, port int, protocol string, action Action) {
	if action != Allow && action != Block {
		fmt.Println("Invalid action: %s. Must be 'allow', or 'block'.", action)
		return
	}

	rule := Rule{
		SourceIP:      sourceIP,
		DestinationIP: destinationIP,
		Port:          port,
		Protocol:      protocol,
		Action:        action,
	}

	rules = append(rules, rule)
	fmt.Println("Added rule: ", rule)
}

// MatchRule checks if a packet matches any rule.
// If no rule matches, it returns the default policy.
func MatchRule(srcIP, dstIP string, port int, protocol string) Action {
	for _, rule := range rules {
		if (rule.SourceIP == srcIP || rule.SourceIP == "*") &&
			(rule.DestinationIP == dstIP || rule.DestinationIP == "*") &&
			(rule.Port == port || rule.Port == 0) &&
			(rule.Protocol == protocol || rule.Protocol == "*") {
			return rule.Action
		}
	}
	return defaultPolicy
}

func RemoveRule(sourceIP, destinationIP string, port int, protocol string) {
	for i, rule := range rules {
		if rule.SourceIP == sourceIP && rule.DestinationIP == destinationIP &&
			(rule.Port == port || port == 0) &&
			(rule.Protocol == protocol || protocol == "*") {
			rules = append(rules[:i], rules[i+1:]...)
			fmt.Println("Removed rule: ", rule)
			return
		}
	}
	fmt.Println("No matching rule found.")
}

func ListRules() {
	fmt.Println("Current Firewall Rules:")
	for _, rule := range rules {
		fmt.Printf("Source IP: %s | Destination IP: %s | Port: %d | Protocol: %s | Action: %s\n",
			rule.SourceIP, rule.DestinationIP, rule.Port, rule.Protocol, rule.Action)
	}
}

func ClearRules() {
	rules = []Rule{}
}

func GetRules() []Rule {
	return rules
}
