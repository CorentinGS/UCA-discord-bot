package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/utils"
)

func VersionCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Version",
						Description: "The current version of the bot is: " + utils.VERSION,
						Author: &discordgo.MessageEmbedAuthor{
							Name:    s.State.User.Username,
							IconURL: s.State.User.AvatarURL(""),
							URL:     utils.GITHUB,
						},
					},
				},
			},
		})
	}
}
