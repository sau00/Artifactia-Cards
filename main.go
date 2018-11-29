package main

import (
	"artifactia-cards/app/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"html/template"
	"io"
	"net/http"
)

const (
	ENV_PROD = 1
	ENV_DEV  = 0
)

var ENV = ENV_PROD

func main() {
	e := echo.New()

	e.Static("/static", "static")
	e.Static("/uploads", "uploads")

	// Database Connection
	db, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	funcMap := template.FuncMap{
		//"displayNewLines": func(s string) template.HTML {
		//	return template.HTML(strings.Replace(s, "\n", "<br />", -1))
		//},
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

	// Routes
	h := &handlers.Handler{
		DB: db,
	}

	e.GET("/", h.FrontendIndexGET)
	e.GET("/parser", h.ParseFunction)

	e.Logger.Fatal(e.Start(":1234"))
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
