package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(generateConfigCmd)
}

var generateConfigCmd = &cobra.Command{
	Use:   "generate-config",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg = &config{
			TracksToSetup: []track{},
		}

		fmt.Printf("generating default config at %s\n", configPath)
		jCfg, err := json.Marshal(cfg)
		if err != nil {
			log.Fatalf("[ERROR] failed to marshal config: %v", err)
		}
		err = os.WriteFile(configPath, jCfg, 0644)
		if err != nil {
			log.Fatalf("[ERROR] failed to write to %s: %v", configPath, err)
		}
	},
}
