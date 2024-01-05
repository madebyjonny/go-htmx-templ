package main

import (
	"context"

	"github.com/labstack/echo/v4"
	templates "github.com/madebyjonny/go-htmx-templ/templates"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        component := templates.TextForm()

        return component.Render(context.Background(), c.Response().Writer)
    })

    e.POST("/submit-form", func(c echo.Context) error {
        name := c.FormValue("name")
        component := templates.Text(name)

        return component.Render(context.Background(), c.Response().Writer)
    })

    e.Logger.Fatal(e.Start(":3000"))

    println("Hello, World!")
}