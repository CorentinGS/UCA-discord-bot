package tag

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/commands/embeds"
	"github.com/corentings/UCA-discord-bot/utils"
)

func AddTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.MessageEmbed {
	options := i.ApplicationCommandData().Options
	if !utils.HasPermissionsAdmin(i.Member) {
		return embeds.CreateForbiddenEmbed(s, i)
	}
	commandOptions := options[0].Options
	key := commandOptions[0].StringValue()
	content := commandOptions[1].StringValue()
	err := addTag(key, content, i.ChannelID, i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err)
	}
	return embeds.CreateSuccessEmbed(s, i, fmt.Sprintf("Tag %s added", key))

}
