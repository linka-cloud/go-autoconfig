package main // import "go-autoconfig"

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	flag "github.com/spf13/pflag"

	"go-autoconfig/config"
	"go-autoconfig/handlers"
	"go-autoconfig/templates"
)

var path = flag.String(
	"config",
	os.Getenv("CONFIG"),
	"enter path to config file",
)

func main() {
	// Parse at first startup
	flag.Parse()
	
	// Read config
	conf, err := config.New(*path)
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(2)
	}
	tmpl := &Template{
		templates: templates.Templates,
	}

	// Init Echo
	e := echo.New()
	e.Renderer = tmpl
	e.HideBanner = true

	// Middleware
	e.Pre(LowerPath)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	h := handlers.Handler{Config: conf}
	e.POST("/autodiscover/autodiscover.xml", h.Outlook)
	e.GET("/mail/config-v1.1.xml", h.Thunderbird)
	e.GET("/.well-known/autoconfig/mail/config-v1.1.xml", h.Thunderbird)
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

func LowerPath(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().URL.Path = strings.ToLower(c.Request().URL.Path)
		next(c)
		return nil
	}
}
