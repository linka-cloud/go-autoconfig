package main // import "go-autoconfig"

import (
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"go-autoconfig/config"
	"go-autoconfig/handlers"
	"go-autoconfig/templates"
)

var path = flag.String(
	"config",
	"",
	"enter path to config file",
)

func main() {
	// Parse at first startup
	flag.Parse()

	// Read config
	conf, err := config.NewConfig(*path)
	if err != nil {
		fmt.Printf("Incorrect path or config itself! See help.\n%s\n", err.Error())
		os.Exit(2)
	}

	tmpl := &Template{
		templates: templates.Templates,
	}

	// Init Echo
	e := echo.New()
	e.Renderer = tmpl

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	h := handlers.Handler{conf}
	e.POST("/autodiscover/autodiscover.xml", h.Outlook)
	e.GET("/mail/config-v1.1.xml", h.Thunderbird)
	e.GET("/email.mobileconfig", h.AppleMail)

	// Start server
	e.Logger.Fatal(e.Start(conf.ServiceAddr))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
