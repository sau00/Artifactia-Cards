package handlers

import (
	"gopkg.in/mgo.v2"
	"github.com/labstack/echo"
	"fmt"
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
