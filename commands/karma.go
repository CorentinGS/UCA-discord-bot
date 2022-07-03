package commands

import (
	"github.com/bwmarrin/discordgo"
)

func KarmaCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		content := ""

		// As you can see, names of subcommands (nested, top-level)
		// and subcommand groups are provided through the arguments.
		switch options[0].Name {
		case "add":
			content = "Add karma command"
			commandOptions := options[0].Options
			for _, opt := range commandOptions {
				if opt.Name == "user-option" {
					content = "go user option"
				}
			}
		case "show":
			commandOptions := options[0].Options
			if len(commandOptions) > 0 {
				for _, opt := range commandOptions {
					if opt.Name == "user-option" {
						content = "go user option"
					}
				}
			} else {
				content = "no user option"
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
