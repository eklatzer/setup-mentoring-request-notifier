package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const defaultGetTracksURL = "https://exercism.org/api/v2/tracks"

var getTracksURL = ""

func init() {
	getTracksCmd.PersistentFlags().StringVar(&getTracksURL, "url", defaultGetTracksURL, "HTTP url to request the track info at")
	cmd.AddCommand(getTracksCmd)
}

var getTracksCmd = &cobra.Command{
	Use:   "get-tracks",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("collecting track infos at %s\n", getTracksURL)
		trackInfo, err := getTracks()
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}
		cfg, err := getConfig(configPath)
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}

		cfg.TracksToSetup = trackInfo.Tracks

		fmt.Printf("writing results to %s\n", configPath)
		jCfg, err := json.Marshal(cfg)
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}
		err = os.WriteFile(configPath, jCfg, 0644)
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}
	},
}

type trackInfo struct {
	Tracks []track `json:"tracks"`
}

type track struct {
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	ChannelID string `json:"channel_id"`
}

func getTracks() (*trackInfo, error) {
	fmt.Printf("requesting tracks at %q\n", getTracksURL)
	resp, err := http.Get(getTracksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to GET track infos %q: %v", getTracksURL, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to GET track infos %q: %v", getTracksURL, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var trackInfo = &trackInfo{}
	err = json.Unmarshal(body, trackInfo)
	if err != nil {
		return nil, err
	}
	return trackInfo, nil
}
