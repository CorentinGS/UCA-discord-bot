package tag

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/commands/embeds"
	"github.com/corentings/UCA-discord-bot/models"
	"github.com/corentings/UCA-discord-bot/utils"
)

func AddTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options
	if !utils.HasPermissionsAdmin(i.Member) {
		return embeds.CreateForbiddenEmbed(s, i), fmt.Errorf("you don't have permissions to add tags")
	}
	imageURL := ""
	commandOptions := options[0].Options
	key := commandOptions[0].StringValue()
	content := commandOptions[1].StringValue()
	if len(commandOptions) > 2 {
		imageURL = commandOptions[2].StringValue()
	}
	err := addTag(key, content, i.ChannelID, i.GuildID, imageURL)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
	}
	return embeds.CreateSuccessEmbed(s, i, fmt.Sprintf("Tag %s added", key)), nil
}
func GetTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options
	commandOptions := options[0].Options
	key := commandOptions[0].StringValue()
	tag, err := models.GetTag(key, i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
	}
	embed := embeds.CreateResponseEmbed(s, i, key, tag.Content)
	if tag.ImageURL != "" {
		embed.Image = &discordgo.MessageEmbedImage{
			URL: tag.ImageURL,
		}
	}
	return embed, nil
}

func DeleteTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options
	if !utils.HasPermissionsAdmin(i.Member) {
		return embeds.CreateForbiddenEmbed(s, i), fmt.Errorf("you don't have permissions to delete tags")
	}
	commandOptions := options[0].Options
	key := commandOptions[0].StringValue()
	err := deleteTag(key, i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
	}
	return embeds.CreateSuccessEmbed(s, i, fmt.Sprintf("Tag %s deleted", key)), nil
}

func HelpTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
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
	return embed, nil
}

func ListTagCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	tags, err := getAllTags(i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
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
			return embeds.CreateErrorEmbed(s, i, err), err
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
	return embed, nil
}
