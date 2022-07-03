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
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"get-version": VersionCommand(),
		"karma":       KarmaCommandHandler(),
		"tag-add":     TagCommandHandler(),
	}
)

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}

func TagCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "The selected message is: " + i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID].Content,
			},
		})
	}
}
