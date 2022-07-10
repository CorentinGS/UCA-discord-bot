package tag

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/commands/embeds"
	"github.com/corentings/UCA-discord-bot/models"
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
func GetTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.MessageEmbed {
	options := i.ApplicationCommandData().Options
	commandOptions := options[0].Options
	key := commandOptions[0].StringValue()
	tag, err := models.GetTag(key, i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err)
	}
	return embeds.CreateResponseEmbed(s, i, key, tag.Content)
}

func DeleteTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.MessageEmbed {
	options := i.ApplicationCommandData().Options
	if !utils.HasPermissionsAdmin(i.Member) {
		return embeds.CreateForbiddenEmbed(s, i)
	}
	commandOptions := options[0].Options
	key := commandOptions[0].StringValue()
	err := deleteTag(key, i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err)
	}
	return embeds.CreateSuccessEmbed(s, i, fmt.Sprintf("Tag %s deleted", key))
}

func HelpTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.MessageEmbed {
	embed := embeds.CreateHelpEmbed(s, i, "tag", "Manage tags")
	commands := []struct {
		Name string
		Desc string
	}{
		{Name: "add",
			Desc: "Add a tag",
		},
		{Name: "get",
			Desc: "Get a tag",
		},
		{Name: "delete",
			Desc: "Delete a tag",
		},
		{Name: "list",
			Desc: "List all tags",
		},
	}
	for _, command := range commands {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  command.Name,
			Value: command.Desc,
		})
	}
	return embed
}

func ListTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.MessageEmbed {
	tags, err := getAllTags(i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err)
	}
	responseContent := "Tags:\n"
	for _, tag := range tags {
		responseContent += fmt.Sprintf("%s\n", tag.Key)
	}

	mappedTags := map[string][]*models.Tag{}
	for _, tag := range tags {
		mappedTags[tag.ChannelID] = append(mappedTags[tag.ChannelID], tag)
	}

	embed := embeds.CreateResponseEmbed(s, i, "Tags", "")

	for channelID, tagSlice := range mappedTags {
		channel, err := s.Channel(channelID)
		if err != nil {
			return embeds.CreateErrorEmbed(s, i, err)
		}
		var result string
		for _, tag := range tagSlice {
			result += fmt.Sprintf("・%s\n", tag.Key)
		}

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("📌 #%s", channel.Name),
			Value: result,
		})
	}
	return embed
}
