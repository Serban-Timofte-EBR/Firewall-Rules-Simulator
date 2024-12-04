package rules

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Config struct {
	DefaultPolicy string `json:"defaultPolicy"`
	Rules         []Rule `json:"rules"`
}

// Load rules and default policy from a configuration file
func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		if err := yaml.NewDecoder(file).Decode(&config); err != nil {
			return fmt.Errorf("failed to parse YAML config file: %v", err)
		}
	} else if strings.HasSuffix(filePath, ".json") {
		if err := json.NewDecoder(file).Decode(&config); err != nil {
			return fmt.Errorf("failed to parse JSON config file: %v", err)
		}
	} else {
		return fmt.Errorf("unsupported config file format: %s", filePath)
	}

	switch config.DefaultPolicy {
	case "allow":
		SetDefaultPolicy(Allow)
	case "block":
		SetDefaultPolicy(Block)
	default:
		return fmt.Errorf("invalid default policy: %s", config.DefaultPolicy)
	}

	for _, rule := range config.Rules {
		AddRule(rule.SourceIP, rule.DestinationIP, rule.Port, rule.Protocol, rule.Action)
	}
	fmt.Println("Configuration loaded successfully.")
	return nil
}
