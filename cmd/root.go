package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const defaultConfigPath = "config.json"

type config struct {
	TracksToSetup []track `json:"track_config"`
}

var cmd = &cobra.Command{}
var configPath string

func init() {
	cmd.PersistentFlags().StringVar(&configPath, "config", defaultConfigPath, "Path and file containing the config")
}

// Execute executes the root command.
func Execute() error {
	return cmd.Execute()
}

func getConfig(path string) (*config, error) {
	jCfg, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s (did you create the default config by running the command 'generate-config'): %w", path, err)
	}
	var cfg = &config{}
	err = json.Unmarshal(jCfg, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
