package models

import (
	"context"
	"github.com/corentings/UCA-discord-bot/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tag struct {
	GuildID   string
	ChannelID string
	Key       string
	Content   string
}

func (tag *Tag) SetTag(GuildID, ChannelID, Key, Content string) {
	tag.Key = Key
	tag.ChannelID = ChannelID
	tag.GuildID = GuildID
	tag.Content = Content
}

func (tag *Tag) DeleteTag() error {
	collection := database.Mg.Db.Collection("tags")
	_, err := collection.DeleteOne(context.TODO(), bson.D{{"key", tag.Key}, {"guildid", tag.GuildID}})
	if err != nil {
		return err
	}
	return nil
}

func (tag *Tag) CreateTag() error {
	collection := database.Mg.Db.Collection("tags")
	_, err := collection.InsertOne(context.TODO(), tag)
	if err != nil {
		return err
	}
	return nil
}

func (tag *Tag) UpdateTag() error {
	collection := database.Mg.Db.Collection("tags")
	_, err := collection.UpdateOne(context.TODO(), bson.D{{"key", tag.Key}, {"guildid", tag.GuildID}}, bson.D{{"$set", bson.D{{"content", tag.Content}}}})
	if err != nil {
		return err
	}
	return nil
}

func GetTag(key, guildID string) (*Tag, error) {
	collection := database.Mg.Db.Collection("tags")
	result := new(Tag)
	err := collection.FindOne(context.TODO(), bson.D{{"key", key}, {"guildid", guildID}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		return nil, err
	}
	return result, nil
}

func GetAllTags(guildID string) ([]*Tag, error) {
	collection := database.Mg.Db.Collection("tags")
	var results []*Tag
	cur, err := collection.Find(context.TODO(), bson.D{{"guildid", guildID}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var result Tag
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func GetAllTagsByChannel(channel string) ([]*Tag, error) {
	collection := database.Mg.Db.Collection("tags")
	var results []*Tag
	cur, err := collection.Find(context.TODO(), bson.D{{"channelid", channel}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var result Tag
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
