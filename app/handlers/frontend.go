package handlers

import (
	"artifactia-cards/app/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	DB *mgo.Session
}

func (h *Handler) Database() (*mgo.Database, *mgo.Session) {
	db := h.DB.Clone()
	return db.DB("artifactia"), db
}

func (h *Handler) FrontendIndexGET(c echo.Context) error {

	db, cn := h.Database()
	defer cn.Close()

	var cards []models.Card

	db.C("cards").Find(bson.M{}).All(&cards)

	return c.Render(200, "frontend/index", struct {
		Cards interface{}
	}{cards})
}