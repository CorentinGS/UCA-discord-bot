package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "get-version",
			Description: "Returns the version of the bot",
		},
		&KarmaCommand,
		&TagCommand,
		&GoogleCommand,
		&DdgCommand,
		&StartpageCommand,
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"get-version": VersionCommand(),
		"karma":       KarmaCommandHandler(),
		"tag":         TagCommandHandler(),
		"google":      GoogleCommandHandler(),
		"ddg":         DdgCommandHandler(),
		"startpage":   StartpageCommandHandler(),
	}
)

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}
