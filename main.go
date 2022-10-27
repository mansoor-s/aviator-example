package main

import (
	"bytes"
	_ "embed"
	"net/http"
	"path/filepath"

	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mansoor-s/aviator"
)

//go:embed exampleTemplate.tmpl
var defaultHTMLTemplate string

var defaultHTMLGenerator = template.Must(template.New("defaultHTML").Parse(defaultHTMLTemplate))

func main() {
	viewsAbsPath, err := filepath.Abs("./views")
	if err != nil {
		panic(err)
	}

	a := aviator.NewAviator(
		aviator.WithViewsPath(viewsAbsPath),
		aviator.WithDevMode(true),
		aviator.WithNumJsVMs(4),
		aviator.WithStaticAssetRoute("/public/assets/"),
		aviator.WithAssetOutputPath("dsfsf"),
	)

	err = a.Init()
	if err != nil {
		panic("Aviator init error: " + err.Error())
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	Root:   "public",
	// 	Browse: true,
	// }))
	//e.Use(middleware.Recover())

	e.GET("/foo", IndexRoute(a))

	e.GET("/gotemplate", func(c echo.Context) error {
		props := struct {
			Myprop         string
			Title          string
			PageHeader     string
			InitCounterVal int
		}{
			Myprop:         "My Prop Value",
			Title:          "This title is passed in from Go",
			PageHeader:     "Page Header from Go",
			InitCounterVal: 42,
		}

		var processed bytes.Buffer
		defaultHTMLGenerator.Execute(&processed, props)

		return c.HTML(http.StatusOK, processed.String())
	})

	e.GET("/public/assets/:asset", func(e echo.Context) error {
		asset, mime, found := a.GetStaticAsset(e.Param("asset"))
		if found {
			return e.Blob(http.StatusOK, mime, asset)
		}
		return e.String(http.StatusNotFound, "404")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
