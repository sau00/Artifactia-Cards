package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"os"
)

type Handler struct {
	DB *mgo.Session
}

func (h *Handler) Database() (*mgo.Database, *mgo.Session) {
	db := h.DB.Clone()
	return db.DB("artifactia"), db
}

func (h *Handler) FrontendIndexGET(c echo.Context) error {
	fmt.Println("Hello world!")
	return nil
}

type CardSet struct {
	CardSet struct {
		SetInfo struct {
			SetId int `json:"set_id"`
			Name  struct {
				Eng string `json:"english"`
				Rus string `json:"russian"`
			} `json:"name"`
		} `json:"set_info"`

		CardList []struct {
			CardId     int    `json:"card_id"`
			BaseCardId int    `json:"base_card_id"`
			CardType   string `json:"card_type"`
			CardName   struct {
				Eng string `json:"english"`
				Rus string `json:"russian"`
			} `json:"card_name"`
			CardText struct {
				Eng string `json:"english"`
				Rus string `json:"russian"`
			} `json:"card_text"`
			MiniImage  string `json:"mini_image"`
			LargeImage struct {
				Eng string `json:"default"`
				Rus string `json:"russian"`
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

func (h *Handler) ParseFunction(c echo.Context) error {
	jsonFile, err := os.Open("main.json")

	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result CardSet

	json.Unmarshal([]byte(byteValue), &result)

	for _, card := range result.CardSet.CardList {

		fmt.Print(card)

	}

	return nil
}
