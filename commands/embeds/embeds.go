package embeds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/utils"
)

func CreateForbiddenEmbed(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.MessageEmbed {
	return createEmbed(s, i, "Forbidden", "⛔ You don't have the required permissions to use this command.", utils.RED)
}

func CreateErrorEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, err error) *discordgo.MessageEmbed {
	return createEmbed(s, i, "Error", fmt.Sprintf("💢 An error occurred️: %s \n", err.Error()), utils.ORANGE)
}

func CreateSuccessEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, message string) *discordgo.MessageEmbed {
	return createEmbed(s, i, "Success", fmt.Sprintf("✅ %s", message), utils.GREEN)
}

func CreateResponseEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, title, content string) *discordgo.MessageEmbed {
	return createEmbed(s, i, title, content, utils.BLUE)
}

func CreateHelpEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, command, description string) *discordgo.MessageEmbed {
	return createEmbed(s, i, fmt.Sprintf("%s help", command), description, utils.PURPLE)
}

func createEmbed(s *discordgo.Session, _ *discordgo.InteractionCreate, title string, description string, color int) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       color,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    s.State.User.Username,
			IconURL: s.State.User.AvatarURL(""),
			URL:     utils.GITHUB,
		},
	}
}
