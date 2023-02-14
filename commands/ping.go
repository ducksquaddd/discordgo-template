package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ducksquaddd/discordgo-template/handlers"
)

func PingCommand() handlers.Command {
	return handlers.Command{
		Name: "Ping",
		Exec: func(s *discordgo.Session, m *discordgo.Message, args *[]string) {
			s.ChannelMessageSend(m.ChannelID, "Pong")
			fmt.Println(args)
		},
	}
}
