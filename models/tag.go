package models

type Tag struct {
	GuildID   string
	ChannelID string
	Key       string
	Content   string
}

func (tag *Tag) SetTag(GuildID string, ChannelID string, Key string, Content string) {
	tag.Key = Key
	tag.ChannelID = ChannelID
	tag.GuildID = GuildID
	tag.Content = Content

}
