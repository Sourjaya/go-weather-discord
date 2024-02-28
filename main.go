package main

import (
	"flag"
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

}
