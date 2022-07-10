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
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "image-option",
					Description: "Image option",
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
		var err error

		switch options[0].Name {
		case "add":
			responseEmbed, err = tag.AddTagCommandHandler(s, i)
		case "get":
			responseEmbed, err = tag.GetTagCommandHandler(s, i)
		case "delete":
			responseEmbed, err = tag.DeleteTagCommandHandler(s, i)
		case "list":
			responseEmbed, err = tag.ListTagCommandHandler(s, i)
		default:
			responseEmbed, err = tag.HelpTagCommandHandler(s, i)
		}

		if err != nil {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{responseEmbed},
					Flags:  discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}

		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{responseEmbed},
			},
		})
	}
}
