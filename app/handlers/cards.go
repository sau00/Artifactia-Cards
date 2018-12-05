package handlers

import (
	"artifactia-cards/app/models"
	"artifactia-cards/app/services"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

func (h *Handler) CardsIndexGET(c echo.Context) error {

	db, cn := h.Database()
	defer cn.Close()

	var cards []models.Card

	db.C("cards").Find(bson.M{}).All(&cards)

	var seo services.Seo

	seo.Title = "Все карты Artifact"
	seo.Description = "Список всех карт из Artifact"

	seo.OG.SiteName = "Artifactia — русскоязычное сообщество игры Artifact"
	seo.OG.URL = "https://artifactia.ru/cards/id/"

	return c.Render(200, "cards/index", struct {
		Seo   interface{}
		Cards interface{}
	}{seo, cards})
}

func (h *Handler) CardsSingleOldGET(c echo.Context) error {

	return nil
}

func (h *Handler) CardGenerateAliases(c echo.Context) error {

	db, cn := h.Database()
	defer cn.Close()

	var cards []models.Card

	db.C("cards").Find(bson.M{}).All(&cards)

	for _, card := range cards {
		db.C("cards").Update(bson.M{"card_id": card.CardId}, bson.M{"$set": bson.M{"seo.alias": services.String2Url(card.CardName.English)}})
	}

	return nil
}

func (h *Handler) CardsSingleGET(c echo.Context) error {
	db, cn := h.Database()
	defer cn.Close()

	var card models.Card

	err := db.C("cards").Find(bson.M{"seo.alias": c.Param("alias")}).One(&card)

	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound}
	}

	var seo services.Seo

	seo.Title = card.CardName.Russian + " карта Artifact"
	seo.Description = "Карта " + card.CardName.Russian + " Artifact"

	seo.OG.Title = card.CardName.Russian + " карта Artifact"
	seo.OG.Description = "Карта " + card.CardName.Russian + " Artifact"
	seo.OG.SiteName = "Artifactia — русскоязычное сообщество игры Artifact"
	seo.OG.Image = "/uploads/cards/large/rus/" + strconv.Itoa(card.CardId) + ".png"
	seo.OG.URL = "https://artifactia.ru/cards/id/" + strconv.Itoa(card.CardId)

	return c.Render(200, "cards/single", struct {
		Seo  interface{}
		Card interface{}
	}{seo, card})
}
