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
		{
			Name:        "add-karma",
			Description: "Add karma to an user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User option",
					Required:    true,
				}},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"get-version": VersionCommand(),
		"add-karma":   AddKarmaCommand(),
	}
)

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}
