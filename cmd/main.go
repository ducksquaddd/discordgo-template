package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ducksquaddd/discordgo-template/commands"
	"github.com/ducksquaddd/discordgo-template/handlers"
	"github.com/ducksquaddd/discordgo-template/helpers"
)

var (
	// Token, Prefix
	BotData helpers.Bot

	// Used to store all the commands
	CmdHandler *handlers.CmdHandler = handlers.New()
)

func main() {
	/* Marshal json and populate BotData */
	helpers.InitJson(&BotData)

	/* Populate the commands */
	CmdHandler.AddCommands(commands.PingCommand())

	var discord, err = discordgo.New(BotData.Token)
	if err != nil {
		panic(err)
	}

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	/* Message Create listener */
	discord.AddHandler(CmdHandler.MessageCreate)

	fmt.Println("Bot online CTRL+C to kill the bot")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	discord.Close()
}
