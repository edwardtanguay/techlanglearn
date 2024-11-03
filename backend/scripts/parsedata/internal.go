package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadConfig() (*Config, error) {
	filename := "../../configs/config-dev.json"

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal JSON into the Config struct
	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config JSON: %w", err)
	}

	return &config, nil
}