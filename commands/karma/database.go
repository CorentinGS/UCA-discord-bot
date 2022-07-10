package karma

import "github.com/corentings/UCA-discord-bot/models"

func addKarma(userID string, guildID string) (*models.Karma, error) {
	karma, err := models.IncreaseKarma(userID, guildID)
	if err != nil {
		return nil, err
	}
	return karma, nil
}
