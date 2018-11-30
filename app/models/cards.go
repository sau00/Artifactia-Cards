package models

import "gopkg.in/mgo.v2/bson"

type Card struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CardId     int           `json:"card_id" bson:"card_id"`
	BaseCardId int           `json:"base_card_id" bson:"base_card_id"`

	CardName struct {
		English string `json:"english" bson:"english"`
		Russian string `json:"russian" bson:"russian"`
	} `json:"card_name" bson:"card_name"`

	CardText struct {
		English string `json:"english" bson:"english"`
		Russian string `json:"russian" bson:"russian"`
	} `json:"card_text" bson:"card_text"`
}