package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/models"
)

var KarmaCommand = discordgo.ApplicationCommand{
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
}

func KarmaCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		content := ""

		// As you can see, names of subcommands (nested, top-level)
		// and subcommand groups are provided through the arguments.
		switch options[0].Name {
		case "add":
			commandOptions := options[0].Options
			for _, opt := range commandOptions {
				if opt.Name == "user-option" {
					if i.Member.User.ID != opt.UserValue(s).ID {
						karma, err := addKarma(opt.UserValue(s).ID, i.GuildID)
						if err != nil {
							content = "Error adding karma + " + err.Error()
						} else {
							content = fmt.Sprintf("Added karma to %s. He now has a karma of %d", opt.UserValue(s).Mention(), karma.Value)
						}
					} else {
						content = "You can't add karma to yourself !"
					}
				}
			}
		case "show":
			commandOptions := options[0].Options
			if len(commandOptions) > 0 {
				for _, opt := range commandOptions {
					if opt.Name == "user-option" {
						karma, err := models.GetKarma(opt.UserValue(s).ID, i.GuildID)
						if err != nil {
							content = "Error while fetching karma: " + err.Error()
						} else {
							content = fmt.Sprintf("%s's karma is : %d", opt.UserValue(s).Username, karma.Value)
						}
					}
				}
			} else {
				karma, err := models.GetKarma(i.Member.User.ID, i.GuildID)
				if err != nil {
					content = "Error while fetching karma: " + err.Error()
				} else {
					content = fmt.Sprintf("Your karma is : %d", karma.Value)
				}
			}

		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	}
}

func addKarma(userID string, guildID string) (*models.Karma, error) {
	karma, err := models.IncreaseKarma(userID, guildID)
	if err != nil {
		return nil, err
	}
	return karma, nil
}
