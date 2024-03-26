/*
Copyright © 2024 Volodymyr Larkin <vlarkin@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

// chatbotCmd represents the chatbot command
var chatbotCmd = &cobra.Command{
	Use:     "chatbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// create telebot object
		chatbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		// exit if failed to connect or print status
		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		} else {
			log.Printf("\nChatbot version %s started", appVersion)
		}
		// process requests
		chatbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload
			// answer on 'hello' and 'joke' requests
			switch payload {
			case "hello":
				// print hello message
				err = m.Send(fmt.Sprintf("Hello I'm ChatBot %s!", appVersion))
			case "joke":
				// print a joke
				joke := getRandomPun()
				err = m.Send(fmt.Sprintf("%s", joke))
			}
			return err

		})

		chatbot.Start()

	},
}

// Get a random pun
func getRandomPun() string {
	puns := []Joke{
		{Setup: "What's orange and sounds like a parrot?", Punchline: "A carrot!"},
		{Setup: "Why couldn’t the leopard play hide and seek?", Punchline: "Because he was always spotted!"},
		{Setup: "Why don't scientists trust atoms?", Punchline: "Because they make up everything!"},
	}
	rand.Seed(int64(len(puns))) // Seed random number generator
	randomIndex := rand.Intn(len(puns))
	return fmt.Sprintf("%s\n%s", puns[randomIndex].Setup, puns[randomIndex].Punchline)
}

func init() {
	rootCmd.AddCommand(chatbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
