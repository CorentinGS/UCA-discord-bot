package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/models"
	"github.com/corentings/UCA-discord-bot/utils"
	"os"
)

var TagCommand = discordgo.ApplicationCommand{
	Name:        "tag",
	Description: "Tag main command",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "add",
			Description: "add tag",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "key-option",
					Description: "Key option",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "content-option",
					Description: "Content option",
					Required:    true,
				},
			},
		},
		{
			Name:        "get",
			Description: "get tag",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "key-option",
					Description: "Key option",
					Required:    true,
				},
			},
		},
		{
			Name:        "delete",
			Description: "delete tag",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "key-option",
					Description: "Key option",
					Required:    true,
				},
			},
		},
		{
			Name:        "list",
			Description: "list tags",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}

func TagCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		responseContent := ""

		// As you can see, names of subcommands (nested, top-level)
		// and subcommand groups are provided through the arguments.
		switch options[0].Name {
		case "add":
			var AdminRole = os.Getenv("ADMIN_ROLE")
			if !utils.ExistsInArray(i.Member.Roles, AdminRole) {
				responseContent = "You are not an admin"
				break
			}
			commandOptions := options[0].Options
			key := commandOptions[0].StringValue()
			content := commandOptions[1].StringValue()
			err := addTag(key, content, i.ChannelID, i.GuildID)
			if err != nil {
				responseContent = err.Error()
			} else {
				responseContent = fmt.Sprintf("Tag %s added", key)
			}
		case "get":
			commandOptions := options[0].Options
			key := commandOptions[0].StringValue()
			tag, err := models.GetTag(key, i.GuildID)
			if err != nil {
				responseContent = err.Error()
			} else {
				responseContent = fmt.Sprintf("Tag %s:\n%s", key, tag.Content)
			}

		case "delete":
			var AdminRole = os.Getenv("ADMIN_ROLE")
			if !utils.ExistsInArray(i.Member.Roles, AdminRole) {
				responseContent = "You are not an admin"
				break
			}
			commandOptions := options[0].Options
			key := commandOptions[0].StringValue()
			err := deleteTag(key, i.GuildID)
			if err != nil {
				responseContent = err.Error()
			} else {
				responseContent = fmt.Sprintf("Tag %s deleted", key)
			}
		case "list":
			tags, err := getAllTags(i.GuildID)
			if err != nil {
				responseContent = err.Error()
			} else {
				responseContent = "Tags:\n"
				for _, tag := range tags {
					responseContent += fmt.Sprintf("%s\n", tag.Key)
				}
			}
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: responseContent,
			},
		})
	}
}

func addTag(key, content, channelID, guildID string) error {
	if key == "" || content == "" {
		return fmt.Errorf("key or content is empty")
	}
	tag, _ := models.GetTag(key, guildID)
	if tag != nil {
		return fmt.Errorf("tag already exists")
	}

	tag = new(models.Tag)
	tag.SetTag(guildID, channelID, key, content)
	err := tag.CreateTag()
	if err != nil {
		return err
	}
	return nil
}

func deleteTag(key, guildID string) error {
	if key == "" {
		return fmt.Errorf("key is empty")
	}
	tag, _ := models.GetTag(key, guildID)
	if tag == nil {
		return fmt.Errorf("tag does not exist")
	}
	err := tag.DeleteTag()
	if err != nil {
		return err
	}
	return nil
}

func getAllTags(guildID string) ([]*models.Tag, error) {
	return models.GetAllTags(guildID)
}

func getAllTagsByChannel(channelID string) ([]*models.Tag, error) {
	return models.GetAllTagsByChannel(channelID)
}
