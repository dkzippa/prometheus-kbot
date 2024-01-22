/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"

	"github.com/spf13/cobra"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

// prometheusKbotCmd represents the prometheusKbot command
var prometheusKbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "starts the application",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot %s started", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("ERROR: please check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Println(m.Message().Payload, m.Text())

			payload := m.Message().Payload
			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm kbot %s!", appVersion))
				if err != nil {
					log.Fatalf("ERROR: can't sent message. %s", err)
					return err
				}
			}

			return err

		})

		//kbot.Handle("/start", func(c telebot.Context) error {
		//	return c.Send("Let's start:")
		//})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(prometheusKbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prometheusKbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prometheusKbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
