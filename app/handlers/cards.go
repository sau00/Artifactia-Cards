package handlers

import (
	"artifactia-cards/app/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func (h *Handler) CardsIndexGET(c echo.Context) error {

	db, cn := h.Database()
	defer cn.Close()

	var cards []models.Card

	db.C("cards").Find(bson.M{}).All(&cards)

	return c.Render(200, "cards/index", struct {
		Cards interface{}
	}{cards})
}

func (h *Handler) CardsSingleGET(c echo.Context) error {
	db, cn := h.Database()
	defer cn.Close()

	var card models.Card

	cardId, _ := strconv.Atoi(c.Param("id"))

	db.C("cards").Find(bson.M{"card_id": cardId}).One(&card)

	return c.Render(200, "cards/single", struct {
		Card interface{}
	}{card})
}
