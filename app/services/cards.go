package services

import (
	"artifactia-cards/app/models"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
	DB *mgo.Session
}

func (s *Service) Database() (*mgo.Database, *mgo.Session) {
	db := s.DB.Clone()
	return db.DB("artifactia"), db
}

type CardSet struct {
	CardSet struct {
		SetInfo struct {
			SetId int `json:"set_id"`
			Name  struct {
				English string `json:"english"`
				Russian string `json:"russian"`
			} `json:"name"`
		} `json:"set_info"`

		CardList []struct {
			CardId     int    `json:"card_id"`
			BaseCardId int    `json:"base_card_id"`
			CardType   string `json:"card_type"`
			CardName   struct {
				English string `json:"english"`
				Russian string `json:"russian"`
			} `json:"card_name"`
			CardText struct {
				English string `json:"english"`
				Russian string `json:"russian"`
			} `json:"card_text"`
			MiniImage  struct {
				Default string `json:"default"`
			} `json:"mini_image"`
			LargeImage struct {
				English string `json:"default"`
				Russian string `json:"russian"`
			} `json:"large_image"`
			IngameImage struct {
				Default string `json:"default"`
			} `ingame_image`
			Rarity      string `json:"rarity"`
			Illustrator string `json:"illustrator"`
			ManaCost    int    `json:"mana_cost"`
			GoldCost    int    `json:"gold_cost"`
		} `json:"card_list"`
	} `json:"card_set"`
}


func (s *Service) ParseCards() {
	jsonFile, err := os.Open("main.json")

	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result CardSet

	json.Unmarshal([]byte(byteValue), &result)

	db, cn := s.Database()
	defer cn.Close()

	for _, card := range result.CardSet.CardList {

		var cardDb models.Card

		cardDb.CardId = card.CardId
		cardDb.BaseCardId = card.BaseCardId

		cardDb.CardName.English = card.CardName.English
		cardDb.CardName.Russian = card.CardName.Russian

		cardDb.CardText.English = card.CardText.English
		cardDb.CardText.Russian = card.CardText.Russian

		fmt.Println(card.CardId)

		_, err := db.C("cards").Upsert(bson.M{"card_id": card.CardId}, cardDb)

		if err != nil {
			fmt.Println(err)
		}

		err = saveImage(card.MiniImage.Default, "uploads/cards/mini/" + strconv.Itoa(card.CardId) + ".png")

		if err != nil {
			fmt.Println(err)
		}
	}
}

func saveImage(url, output string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}