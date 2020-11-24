package rekordbox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	Options [][]string `json:"options"`
}

func parseAgentConfig() (*config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("os user home dir failed: %w", err)
	}

	optionsPath := getAgentOptionsPath(home)

	data, err := ioutil.ReadFile(optionsPath)
	if err != nil {
		return nil, fmt.Errorf("ioutil read file failed: %w", err)
	}

	var cfg config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return &cfg, nil
}
