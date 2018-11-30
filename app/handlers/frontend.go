package handlers

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
)

type Handler struct {
	DB *mgo.Session
}

func (h *Handler) Database() (*mgo.Database, *mgo.Session) {
	db := h.DB.Clone()
	return db.DB("artifactia"), db
}

func (h *Handler) FrontendIndexGET(c echo.Context) error {

	return c.Render(200, "frontend/index", nil)
}
