package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type weatherData struct {
	Location struct {
		Name      string `json:"name"`
		Region    string `json:"region"`
		Country   string `json:"country"`
		Localtime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Celsius  float64 `json:"temp_c"`
		Windkph  float64 `json:"wind_kph"`
		Humidity float64 `json:"humidity"`
	} `json:"current"`
}

func query(city, token string) (weatherData, error) {
	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + city

	req, _ := http.NewRequest("GET", url, http.NoBody)
	req.Header.Add("X-RapidAPI-Key", token)
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return weatherData{}, err
	}

	defer res.Body.Close()

	var d weatherData

	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}
func getWeather(message, weatherAPIToken string) *discordgo.MessageSend {
	// Build Discord embed response
	str := strings.Split(message, " ")
	if str[0] == "!weather" {
		loc := strings.Join(str[1:], " ")
		data, err := query(loc, weatherAPIToken)

		if err != nil {
			log.Fatalf("Error fetching weather data: %v", err)

			return &discordgo.MessageSend{
				Content: "Sorry, there was an error trying to get the weather",
			}
		}

		celsius := fmt.Sprintf("%0.2f", data.Current.Celsius)
		humidity := fmt.Sprintf("%0.2f", data.Current.Humidity)
		wind := fmt.Sprintf("%0.2f", data.Current.Windkph)

		embed := &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       "Current Weather",
				Description: "Weather for " + data.Location.Name,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Region",
						Value:  data.Location.Region,
						Inline: true,
					},
					{
						Name:   "Country",
						Value:  data.Location.Country,
						Inline: true,
					},
					{
						Name:   "Local Time",
						Value:  data.Location.Localtime,
						Inline: true,
					},
					{
						Name:   "Country",
						Value:  data.Location.Country,
						Inline: true,
					},
					{
						Name:   "Conditions",
						Value:  data.Current.Condition.Text,
						Inline: true,
					},
					{
						Name:   "Temperature",
						Value:  celsius + "°C",
						Inline: true,
					},
					{
						Name:   "Humidity",
						Value:  humidity + "%",
						Inline: true,
					},
					{
						Name:   "Wind",
						Value:  wind + " mph",
						Inline: true,
					},
				},
			},
			},
		}

		return embed
	}

	return &discordgo.MessageSend{
		Content: "Sorry, please write the command at the beginning of your message",
	}
}
