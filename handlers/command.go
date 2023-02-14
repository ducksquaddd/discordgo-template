package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// ! Creates the commands interface
func New() *CmdHandler {
	return &CmdHandler{
		Commands: map[string]Command{},
	}
}

// ! The guide every command must follow
type Command struct {
	Name string
	Exec func(s *discordgo.Session, m *discordgo.Message, args *[]string)
}

// ! So You can store mutiple commands
type CmdHandler struct {
	Commands map[string]Command
}

func (h *CmdHandler) AddCommands(commands ...Command) {
	for _, command := range commands {
		h.Commands[command.Name] = command
	}
}

// ! Wish i could throw this into a seperate file but it needs the handler struct

func (h *CmdHandler) MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	var content string = message.Content
	var prefix string = "!!"

	if message.Author.Bot {
		return
	}

	if !strings.HasPrefix(content, prefix) {
		return
	}

	// ! delete the prefix, and split every argument into a slice
	var args []string = strings.Fields(strings.TrimSpace(content[len(prefix):]))

	var cmd, ok = h.Commands[args[0]]

	// ! Should only run if they did the prefix without a valid command
	if !ok {
		return
	}

	cmd.Exec(session, message.Message, &args)
}
