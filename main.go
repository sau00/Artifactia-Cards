package main

import (
	"artifactia-cards/app/handlers"
	"artifactia-cards/app/services"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"html/template"
	"io"
	"net/http"
	"os"
)

const (
	ENV_PROD = 1
	ENV_DEV  = 0
)

var ENV = ENV_PROD

func main() {

	// Database Connection
	db, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println(err)
	}

	if len(os.Args) > 1 {

		args := os.Args[1:]

		switch args[0] {

		case "--cards":
			s := &services.Service{
				DB: db,
			}

			s.ParseCards()

		default:
			fmt.Println("Can't enter the matrix")
		}
	} else {
		e := echo.New()

		e.Static("/static", "static")
		e.Static("/uploads", "uploads")

		funcMap := template.FuncMap{
			"html": func(s string) template.HTML {
				return template.HTML(s)
			},
		}

		// Template Engine
		t := &Template{
			templates: template.Must(template.New("lks").Funcs(funcMap).ParseGlob("app/views/*/*.html")),
		}

		if ENV == ENV_PROD {
			e.HTTPErrorHandler = customHTTPErrorHandler
		} else if ENV == ENV_DEV {
			e.Debug = true
		}

		// Configuration
		e.Renderer = t
		e.Logger.SetLevel(log.ERROR)

		h := &handlers.Handler{
			DB: db,
		}

		// Routes
		e.GET("/", h.CardsIndexGET)
		e.GET("/card/:id", h.CardsSingleGET)

		e.Logger.Fatal(e.Start(":1233"))
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	type Data struct {
		Data interface{}
	}

	return t.templates.ExecuteTemplate(w, name, Data{data})
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	err = c.Render(code, "layouts/error", struct {
		Code int
	}{code})
}
