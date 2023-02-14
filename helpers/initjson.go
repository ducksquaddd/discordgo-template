package helpers

import (
	"encoding/json"
	"os"
)

type Bot struct {
	Token string `json:"token"`
}

func InitJson(BotData *Bot) {
	botData, err := os.ReadFile("./config/botdata.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(botData, &BotData)
}
