package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken        string
	WeatherAPIToken string
)

func Start() {
	// Create a new Discord session using the provided bot token.
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}
	// Register the mSend func as a callback for MessageCreate events.
	discord.AddHandler(mSend)
	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		log.Fatalf("error opening websocket connection: %v", err)
		return
	}
	defer discord.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running")

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
func mSend(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot message
	if message.Author.ID == discord.State.User.ID {
		return
	}

	helpMessage := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Help",
			Description: "Help for Weather Bot",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Command Syntax",
					Value:  "!weather <location name>",
					Inline: true,
				},
				{
					Name:   "Command Example 1",
					Value:  "!weather kolkata",
					Inline: true,
				},
				{
					Name:   "Command Example 2",
					Value:  "!weather Rio De Janeiro",
					Inline: true,
				},
			},
		},
		},
	}
	// Respond to messages

	if message.Content == "!weatherhelp" {
		if _, err := discord.ChannelMessageSendComplex(message.ChannelID, helpMessage); err != nil {
			log.Fatalf("Error while sending message over channel: %v", err)
		}
	} else if strings.Contains(message.Content, "!weather") {
		cWeather := getWeather(message.Content, WeatherAPIToken)
		if _, err := discord.ChannelMessageSendComplex(message.ChannelID, cWeather); err != nil {
			log.Fatalf("Error while sending message over channel: %v", err)
		}
	}
}
