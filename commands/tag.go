package commands

import "github.com/bwmarrin/discordgo"

var TagCommand = discordgo.ApplicationCommand{
	Name:        "tag",
	Description: "Tag main command",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "add",
			Description: "add tag",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "key-option",
					Description: "Key option",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "content-option",
					Description: "Content option",
					Required:    true,
				},
			},
		},
		{
			Name:        "show",
			Description: "show tag",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "key-option",
					Description: "Key option",
					Required:    true,
				},
			},
		},
		{
			Name:        "list",
			Description: "list tags",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}
