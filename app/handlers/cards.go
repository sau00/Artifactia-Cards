package handlers

import (
	"artifactia-cards/app/models"
	"artifactia-cards/app/services"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) CardsIndexGET(c echo.Context) error {

	db, cn := h.Database()
	defer cn.Close()

	var cards []models.Card

	filter := c.Param("filter")

	if filter == "hero" || filter == "spell" || filter == "creep" || filter == "improvement" || filter == "item" {
		db.C("cards").Find(bson.M{"card_type": strings.Title(filter)}).All(&cards)
	} else if filter == "uncommon" || filter == "common" || filter == "rare" {
		db.C("cards").Find(bson.M{"rarity": strings.Title(filter)}).All(&cards)
	} else if filter == "black" || filter == "green" || filter == "red" || filter == "blue" {
		db.C("cards").Find(bson.M{"color": filter}).All(&cards)
	} else if filter == "" {
		db.C("cards").Find(bson.M{}).All(&cards)
	} else {
		return &echo.HTTPError{Code: http.StatusNotFound}
	}

	var seo services.Seo

	// Refactor this shit for the first in the next patch!!!!!!!!!!!!!!!!!!!!!
	if filter == "" {
		seo.Title = "Все карты Artifact"
		seo.Description = "Список всех карт из Artifact"
	}

	if filter == "hero" {
		seo.Title = "Все карты героев"
		seo.Description = "Список всех карт героев из Artifact"
	}

	if filter == "spell" {
		seo.Title = "Все карты заклинаний Artifact"
		seo.Description = "Список всех карт заклинаний из Artifact"
	}

	if filter == "creep" {
		seo.Title = "Все карты крипов Artifact"
		seo.Description = "Список всех карт крипов из Artifact"
	}

	if filter == "improvement" {
		seo.Title = "Все карты улучшений Artifact"
		seo.Description = "Список всех карт улучшений из Artifact"
	}

	if filter == "item" {
		seo.Title = "Все карты предметов Artifact"
		seo.Description = "Список всех карт предметов из Artifact"
	}

	// Rarity
	if filter == "common" {
		seo.Title = "Все обычные карты"
		seo.Description = "Список всех обычных карт из Artifact"
	}

	if filter == "uncommon" {
		seo.Title = "Все необычные карты"
		seo.Description = "Список всех необычных карт из Artifact"
	}

	if filter == "rare" {
		seo.Title = "Все редкие карты"
		seo.Description = "Список всех редких карт из Artifact"
	}

	if filter == "blue" {
		seo.Title = "Все синие карты Artifact"
		seo.Description = "Список всех синих карт из Artifact"
	}
	if filter == "green" {
		seo.Title = "Все зеленые карты Artifact"
		seo.Description = "Список всех зеленых карт из Artifact"
	}
	if filter == "black" {
		seo.Title = "Все черные карты Artifact"
		seo.Description = "Список всех черных карт из Artifact"
	}
	if filter == "red" {
		seo.Title = "Все красные карты Artifact"
		seo.Description = "Список всех красных карт из Artifact"
	}

	seo.OG.SiteName = "Artifactia — русскоязычное сообщество игры Artifact"
	seo.OG.URL = "https://artifactia.ru/cards"

	return c.Render(200, "cards/index", struct {
		Seo   interface{}
		Cards interface{}
		Filter string
	}{seo, cards, filter})
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

	var seo services.Seo

	err := db.C("cards").Find(bson.M{"seo.alias": c.Param("alias")}).One(&card)

	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound}
	}

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

func (h *Handler) CardsSingleOldGET(c echo.Context) error {

	db, cn := h.Database()
	defer cn.Close()

	var card models.Card

	cardId, err := strconv.Atoi(c.Param("id"))

	err = db.C("cards").Find(bson.M{"card_id": cardId}).One(&card)

	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound}
	}

	return c.Redirect(http.StatusMovedPermanently, "/cards/" + card.SEO.Alias)
}