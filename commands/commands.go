package commands

import (
	"github.com/bwmarrin/discordgo"
)

const Version = "v0.0.0-alpha"

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "get-version",
			Description: "Returns the version of the bot",
		},
		&KarmaCommand,
		{
			Name: "tag-add",
			Type: discordgo.MessageApplicationCommand,
		},
		&TagCommand,
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"get-version": VersionCommand(),
		"karma":       KarmaCommandHandler(),
		"tag":         TagCommandHandler(),
	}
)

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}
