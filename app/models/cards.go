package models

import "gopkg.in/mgo.v2/bson"

type Card struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`

	SEO struct {
		Alias string `json:"alias" bson:"alias"`
	}

	CardId     int `json:"card_id" bson:"card_id"`
	BaseCardId int `json:"base_card_id" bson:"base_card_id"`

	CardType  string `json:"card_type" bson:"card_type"`
	Rarity    string `json:"rarity" bson:"rarity"`
	Color     string `json:"color" bson:"color"`
	ManaCost  int    `json:"mana_cost" bson:"mana_cost"`
	GoldCost  int    `json:"gold_cost" bson:"gold_cost"`
	HitPoints int    `json:"hit_points" bson:"hit_points"`
	Attack    int    `json:"attack" bson:"attack"`

	// Lore
	Illustrator string `json:"illustrator" bson:"illustrator"`

	CardName struct {
		English string `json:"english" bson:"english"`
		Russian string `json:"russian" bson:"russian"`
	} `json:"card_name" bson:"card_name"`

	CardText struct {
		English string `json:"english" bson:"english"`
		Russian string `json:"russian" bson:"russian"`
	} `json:"card_text" bson:"card_text"`
}
