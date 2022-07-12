package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

const exercismColorCode = "#604FCD"

type trackConfig struct {
	ThreadTS  string `json:"thread_ts"`
	ChannelID string `json:"channel_id"`
}

var slackToken = ""
var slackDelay = 0
var defaultSlackChannelID = ""

func init() {
	setupSlackCmd.PersistentFlags().StringVar(&slackToken, "slack-token", "", "Token used for auth at Slack")
	setupSlackCmd.PersistentFlags().IntVar(&slackDelay, "slack-delay", 500, "Delay in ms after sending a message to Slack (rate limiting)")
	setupSlackCmd.PersistentFlags().StringVar(&defaultSlackChannelID, "default-channel-id", "", "Default value used as channel id if not set")
	cmd.AddCommand(setupSlackCmd)
}

var setupSlackCmd = &cobra.Command{
	Use:   "setup-slack",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting setup of Slack")
		cfg, err := getConfig(configPath)
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}

		slackTrackConfig := map[string]trackConfig{}

		fmt.Println("sending messages to start thread to Slack")

		client := slack.New(slackToken)
		for _, trackInfo := range cfg.TracksToSetup {
			if trackInfo.ChannelID == "" {
				trackInfo.ChannelID = defaultSlackChannelID
			}
			fmt.Printf("> starting thread for %s in channel %s\n", trackInfo.Title, trackInfo.ChannelID)
			_, threadTS, _, err := client.SendMessage(trackInfo.ChannelID, slack.MsgOptionAttachments(slack.Attachment{Text: fmt.Sprintf("Thread for `%s`", trackInfo.Title), Color: exercismColorCode}))
			if err != nil {
				log.Fatalf("[ERROR] %v", err)
			}

			slackTrackConfig[trackInfo.Slug] = trackConfig{
				ThreadTS:  threadTS,
				ChannelID: trackInfo.ChannelID,
			}
			fmt.Printf(">> thread timestamp: %s\n", threadTS)
			time.Sleep(time.Millisecond * time.Duration(slackDelay))
		}

		fmt.Println("finished sending messages")

		jSlackTrackConfig, err := json.Marshal(slackTrackConfig)
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}

		fmt.Println("output for 'track_config' in config of mentoring request notifier:")
		fmt.Println(string(jSlackTrackConfig))
	},
}
