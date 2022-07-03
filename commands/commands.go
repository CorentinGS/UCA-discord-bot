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
			Name:        "karma",
			Description: "Karma main command",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "add",
					Description: "add karma",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionUser,
							Name:        "user-option",
							Description: "User option",
							Required:    true,
						},
					},
				},
				{
					Name:        "show",
					Description: "show karma",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionUser,
							Name:        "user-option",
							Description: "User option",
							Required:    false,
						},
					},
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"get-version": VersionCommand(),
		"karma":       KarmaCommand(),
	}
)

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}
