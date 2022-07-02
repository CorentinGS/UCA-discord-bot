package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func AddKarmaCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		// Or convert the slice into a map
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		// This example stores the provided arguments in an []interface{}
		// which will be used to format the bot's response
		margs := make([]interface{}, 0, len(options))
		msgformat := "You learned how to use command options! " +
			"Take a look at the value(s) you entered:\n"

		if opt, ok := optionMap["user-option"]; ok {
			margs = append(margs, opt.UserValue(nil).ID)
			msgformat += "> user-option: <@%s>\n"
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					msgformat,
					margs...,
				),
			},
		})
	}
}
