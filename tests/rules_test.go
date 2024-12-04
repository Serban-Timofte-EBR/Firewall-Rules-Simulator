package tests

import (
	"Firewall-Rules-Simulator/rules"
	"testing"
)

func TestAddRule(t *testing.T) {
	rules.ClearRules()

	rules.AddRule("192.168.1.1", "*", 80, "TCP", rules.Allow)

	currentRules := rules.GetRules()
	if len(currentRules) != 1 {
		t.Errorf("Expected 1 rule, got %d", len(currentRules))
	}

	expected := rules.Rule{
		SourceIP:      "192.168.1.1",
		DestinationIP: "*",
		Port:          80,
		Protocol:      "TCP",
		Action:        rules.Allow,
	}

	if currentRules[0] != expected {
		t.Errorf("Expected rule %v, got %v", expected, currentRules[0])
	}
}

func TestMatchRule(t *testing.T) {
	rules.ClearRules()

	rules.AddRule("192.168.1.1", "*", 80, "TCP", rules.Allow)
	rules.AddRule("*", "192.168.1.101", 22, "TCP", rules.Block)
	rules.AddRule("10.0.0.1", "10.0.0.2", 0, "*", rules.Allow)

	testCases := []struct {
		srcIP    string
		dstIP    string
		port     int
		protocol string
		expected rules.Action
	}{
		{"192.168.1.1", "192.168.1.2", 80, "TCP", rules.Allow},
		{"192.168.1.2", "192.168.1.101", 22, "TCP", rules.Block},
		{"10.0.0.1", "10.0.0.2", 443, "UDP", rules.Allow},
		{"192.168.1.3", "192.168.1.102", 443, "TCP", rules.Block},
	}

	for _, tc := range testCases {
		result := rules.MatchRule(tc.srcIP, tc.dstIP, tc.port, tc.protocol)
		if result != tc.expected {
			t.Errorf("Expected %s for packet from %s to %s, got %s", tc.expected, tc.srcIP, tc.dstIP, result)
		}
	}
}

func TestDefaultPolicy(t *testing.T) {
	rules.ClearRules()

	rules.SetDefaultPolicy(rules.Block)
	action := rules.MatchRule("192.168.1.3", "192.168.1.102", 443, "TCP")
	if action != rules.Block {
		t.Errorf("Expected default policy 'block', got '%v'", action)
	}

	rules.SetDefaultPolicy(rules.Allow)
	action = rules.MatchRule("192.168.1.3", "192.168.1.102", 443, "TCP")
	if action != rules.Allow {
		t.Errorf("Expected default policy 'allow', got '%v'", action)
	}
}

func TestRemoveRule(t *testing.T) {
	rules.ClearRules()

	rules.AddRule("192.168.1.1", "*", 80, "TCP", rules.Allow)
	rules.AddRule("*", "192.168.1.101", 22, "TCP", rules.Block)

	rules.RemoveRule("192.168.1.1", "*", 80, "TCP")

	currentRules := rules.GetRules()
	if len(currentRules) != 1 {
		t.Errorf("Expected 1 rule after removal, got %d", len(currentRules))
	}

	expected := rules.Rule{
		SourceIP:      "*",
		DestinationIP: "192.168.1.101",
		Port:          22,
		Protocol:      "TCP",
		Action:        rules.Block,
	}

	if currentRules[0] != expected {
		t.Errorf("Expected rule %v, got %v", expected, currentRules[0])
	}
}
