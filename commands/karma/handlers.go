package karma

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/commands/embeds"
	"github.com/corentings/UCA-discord-bot/models"
)

func AddKarmaCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options

	commandOptions := options[0].Options
	for _, opt := range commandOptions {
		if opt.Name == "user-option" {
			if i.Member.User.ID == opt.UserValue(s).ID {
				return embeds.CreateForbiddenEmbed(s, i), fmt.Errorf("you can't add karma to yourself")
			}

			karma, err := addKarma(opt.UserValue(s).ID, i.GuildID)
			if err != nil {
				return embeds.CreateErrorEmbed(s, i, err), err
			}

			return embeds.CreateSuccessEmbed(s, i, fmt.Sprintf("Added karma to %s. He now has a karma of %d", opt.UserValue(s).Mention(), karma.Value)), nil
		}
	}
	return embeds.CreateErrorEmbed(s, i, fmt.Errorf("an error occurred while processing the command")), fmt.Errorf("an error occurred while processing the command")
}

func ShowKarmaCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	options := i.ApplicationCommandData().Options

	commandOptions := options[0].Options
	var user *discordgo.User
	if len(commandOptions) > 0 {
		if commandOptions[0].Name == "user-option" {
			user = commandOptions[0].UserValue(s)
		} else {
			return embeds.CreateErrorEmbed(s, i, fmt.Errorf("an error occurred while processing the command")), fmt.Errorf("an error occurred while processing the command")
		}
	} else {
		user = i.Member.User
	}
	karma, err := models.GetKarma(user.ID, i.GuildID)
	if err != nil {
		return embeds.CreateErrorEmbed(s, i, err), err
	}
	return embeds.CreateResponseEmbed(s, i, user.Username, fmt.Sprintf("%d", karma.Value)), nil
}

func HelpKarmaCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, error) {
	embed := embeds.CreateHelpEmbed(s, i, "karma", "Manage karma")
	commands := []struct {
		Name string
		Desc string
	}{
		{
			Name: "add",
			Desc: "Add karma to a user",
		},
		{
			Name: "show",
			Desc: "Show karma of a user",
		},
	}
	for _, cmd := range commands {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  cmd.Name,
			Value: cmd.Desc,
		})
	}
	return embed, nil
}
