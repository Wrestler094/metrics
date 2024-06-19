package configs

import "strings"

type AgentConfig struct {
	ServerAddress  string `env:"ADDRESS"`
	PollInterval   int64  `env:"REPORT_INTERVAL"`
	ReportInterval int64  `env:"POLL_INTERVAL"`
}

func ValidateAgentConfig(cfg *AgentConfig) {
	if cfg.PollInterval < 1 {
		cfg.PollInterval = 2
	}

	if cfg.ReportInterval < 1 {
		cfg.ReportInterval = 10
	}

	if !(strings.HasPrefix(cfg.ServerAddress, "http://")) {
		cfg.ServerAddress = "http://" + cfg.ServerAddress
	}
}
