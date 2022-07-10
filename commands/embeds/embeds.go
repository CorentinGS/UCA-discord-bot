package embeds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/utils"
)

const GREEN = 0x00ff00
const RED = 0xff0000
const BLUE = 0x0000ff
const ORANGE = 0xffa500

func CreateForbiddenEmbed(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.MessageEmbed {
	return createEmbed(s, i, "Forbidden", "⛔ You don't have the required permissions to use this command.", RED)
}

func CreateErrorEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, err error) *discordgo.MessageEmbed {
	return createEmbed(s, i, "Error", fmt.Sprintf("💢 An error occurred️: %s \n", err.Error()), ORANGE)
}

func CreateSuccessEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, message string) *discordgo.MessageEmbed {
	return createEmbed(s, i, "Success", fmt.Sprintf("✅ %s", message), GREEN)
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
