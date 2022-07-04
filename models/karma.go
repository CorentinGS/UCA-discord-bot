package models

import (
	"context"
	"github.com/corentings/UCA-discord-bot/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Karma struct {
	UserID  string `json:"userid"`
	GuildID string `json:"guildid"`
	Value   uint   `json:"value"`
}

func (karma *Karma) CreateKarma() error {
	collection := database.Mg.Db.Collection("karma")

	_, err := collection.InsertOne(context.TODO(), karma)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (karma *Karma) UpdateKarma() error {
	collection := database.Mg.Db.Collection("karma")

	_, err := collection.UpdateOne(context.TODO(), bson.D{{"userid", karma.UserID}, {"guildid", karma.GuildID}}, bson.D{{"$set", bson.D{{"value", karma.Value}}}})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (karma *Karma) SetKarma(userID, guildID string, value uint) {
	karma.UserID = userID
	karma.GuildID = guildID
	karma.Value = value
}

func (karma *Karma) AddKarma(amount uint) {
	karma.Value += amount
}

func GetKarma(userID, guildID string) (*Karma, error) {
	collection := database.Mg.Db.Collection("karma")
	result := new(Karma)
	err := collection.FindOne(context.TODO(), bson.D{{"userid", userID}, {"guildid", guildID}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			result.SetKarma(userID, guildID, 0)
			err = result.CreateKarma()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func IncreaseKarma(userID, guildID string) (*Karma, error) {
	collection := database.Mg.Db.Collection("karma")
	result := new(Karma)
	err := collection.FindOne(context.TODO(), bson.D{{"userid", userID}, {"guildid", guildID}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			result.SetKarma(userID, guildID, 1)

			err = result.CreateKarma()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		result.AddKarma(1)
		err := result.UpdateKarma()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
