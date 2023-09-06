package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/utils"
	"strings"
)

var GoogleCommand = discordgo.ApplicationCommand{
	Name:        "google",
	Description: "Google main command",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "search",
			Description: "Search",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

var DdgCommand = discordgo.ApplicationCommand{
	Name:        "ddg",
	Description: "Duckduckgo main command",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "search",
			Description: "Search",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

var StartpageCommand = discordgo.ApplicationCommand{
	Name:        "startpage",
	Description: "Startpage main command",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "search",
			Description: "Search",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

func GoogleCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		var embed *discordgo.MessageEmbed

		search := options[0].StringValue()

		// replace space with +
		search = strings.Replace(search, " ", "+", -1)

		embed = &discordgo.MessageEmbed{
			Title:       "Google Search",
			Description: "Votre recherche est prête:  **[Voir le resultat](http://lmgtfy2.com/?q=" + search + (")**"),
			Color:       utils.BLUE,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    s.State.User.Username,
				IconURL: s.State.User.AvatarURL(""),
				URL:     utils.GITHUB,
			},
		}

		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
			},
		})
	}
}

func DdgCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		var embed *discordgo.MessageEmbed

		search := options[0].StringValue()
		// replace space with +
		search = strings.Replace(search, " ", "+", -1)

		embed = &discordgo.MessageEmbed{
			Title:       "Duckduckgo",
			Description: "Votre recherche est prête:  **[Voir le resultat](http://lmgtfy2.com/?s=d&q=" + search + (")**"),
			Color:       utils.ORANGE,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    s.State.User.Username,
				IconURL: s.State.User.AvatarURL(""),
				URL:     utils.GITHUB,
			},
		}

		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
			},
		})
	}
}

func StartpageCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		var embed *discordgo.MessageEmbed

		search := options[0].StringValue()
		// replace space with +
		search = strings.Replace(search, " ", "+", -1)

		embed = &discordgo.MessageEmbed{
			Title:       "Startpage",
			Description: "Votre recherche est prête:  **[Voir le resultat](https://lmsptfy.com/?q=" + search + (")**"),
			Color:       utils.PURPLE,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    s.State.User.Username,
				IconURL: s.State.User.AvatarURL(""),
				URL:     utils.GITHUB,
			},
		}

		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
			},
		})
	}
}
