package tag

import (
	"fmt"
	"github.com/corentings/UCA-discord-bot/models"
)

func addTag(key, content, channelID, guildID, imageURL string) error {
	if key == "" || content == "" {
		return fmt.Errorf("key or content is empty")
	}
	tag, _ := models.GetTag(key, guildID)
	if tag != nil {
		return fmt.Errorf("tag already exists")
	}

	tag = new(models.Tag)
	tag.SetTag(guildID, channelID, key, content, imageURL)
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
