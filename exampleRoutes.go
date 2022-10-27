package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mansoor-s/aviator"
)

func IndexRoute(a *aviator.Aviator) func(c echo.Context) error {
	return func(c echo.Context) error {
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
		rendered, err := a.Render(c.Request().Context(), "Index.svelte", props)
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}

		return c.HTML(http.StatusOK, rendered)
	}
}
