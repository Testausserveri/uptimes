package front

import (
	"fmt"
	"html/template"
	"log"

	"github.com/Testausserveri/uptimes/core"
	"github.com/Testausserveri/uptimes/storage"
	"github.com/labstack/echo/v4"
)

func StatusGroupHandler(group core.StatusGroup) echo.HandlerFunc {
	return func(c echo.Context) error {
		templatePath := fmt.Sprintf("public/%s", group.Config.TemplateName)
		t, err := template.ParseFiles(templatePath)
		if err != nil {
			log.Println(err)
			return err
		}

		return t.Execute(c.Response().Writer, struct{ Storage storage.Storage }{
			*group.Storage,
		})
	}
}

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Static("/assets", "public/assets")
	return e
}
