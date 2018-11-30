package handlers

import (
	"fmt"
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
	fmt.Println("Hello world!")
	return nil
}
