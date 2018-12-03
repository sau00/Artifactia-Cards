package handlers

import (
	"artifactia-cards/app/models"
	"artifactia-cards/app/services"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
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
		Seo interface{}
		Cards interface{}
	}{seo, cards})
}

func (h *Handler) CardsSingleGET(c echo.Context) error {
	db, cn := h.Database()
	defer cn.Close()

	var card models.Card

	cardId, _ := strconv.Atoi(c.Param("id"))

	db.C("cards").Find(bson.M{"card_id": cardId}).One(&card)


	var seo services.Seo

	seo.Title = card.CardName.Russian + " карта Artifact"
	seo.Description = "Карта " + card.CardName.Russian + " Artifact"


	seo.OG.Title = card.CardName.Russian + " карта Artifact"
	seo.OG.Description = "Карта " + card.CardName.Russian + " Artifact"
	seo.OG.SiteName = "Artifactia — русскоязычное сообщество игры Artifact"
	seo.OG.Image = "/uploads/cards/large/rus/" + strconv.Itoa(card.CardId) + ".png"
	seo.OG.URL = "https://artifactia.ru/cards/id/" + strconv.Itoa(card.CardId)

	return c.Render(200, "cards/single", struct {
		Seo interface{}
		Card interface{}
	}{seo, card})
}