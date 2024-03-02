package main

import (
	"flag"

	"github.com/Sourjaya/go-weather-discord/bot"
)

// Variables used for command line parameters
var (
	BotToken        string
	WeatherAPIToken string
)

func init() {
	flag.StringVar(&BotToken, "b", "", "Bot Token")
	flag.StringVar(&WeatherAPIToken, "w", "", "Weather API Token")
	flag.Parse()
}

func main() {
	// Save Bot and API keys and start bot
	bot.BotToken = BotToken
	bot.WeatherAPIToken = WeatherAPIToken
	bot.Start()
}
