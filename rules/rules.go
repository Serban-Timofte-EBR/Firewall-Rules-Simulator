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
	Action        Action
}

var rules []Rule

// AddRule adds a new rule to the list
func AddRule(sourceIP, destinationIP string, port int, action Action) {
	if action != Allow && action != Block {
		fmt.Println("Invalid action: %s. Must be 'allow', or 'block'.", action)
		return
	}

	rule := Rule{
		SourceIP:      sourceIP,
		DestinationIP: destinationIP,
		Port:          port,
		Action:        action,
	}

	rules = append(rules, rule)
	fmt.Println("Added rule: ", rule)
}

// MatchRule checks if a packet matches any rule
func MatchRule(sourceIP, destinationIP string, port int) Action {
	for _, rule := range rules {
		if rule.SourceIP == sourceIP && rule.DestinationIP == destinationIP && rule.Port == port {
			return rule.Action
		}
	}
	return "" // No match
}
