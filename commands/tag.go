package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/commands/tag"
)

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
			Name:        "get",
			Description: "get tag",
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
			Name:        "delete",
			Description: "delete tag",
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
		{
			Name:        "help",
			Description: "help",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}

func TagCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		var responseEmbed *discordgo.MessageEmbed

		switch options[0].Name {
		case "add":
			responseEmbed = tag.AddTagCommandHandler(s, i)
		case "get":
			responseEmbed = tag.GetTagCommandHandler(s, i)
		case "delete":
			responseEmbed = tag.DeleteTagCommandHandler(s, i)
		case "list":
			responseEmbed = tag.ListTagCommandHandler(s, i)
		default:
			responseEmbed = tag.HelpTagCommandHandler(s, i)
		}

		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{responseEmbed},
			},
		})
	}
}
