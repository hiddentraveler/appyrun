package handlers

import (
	"github.com/labstack/echo/v4"
	"html/template"
)

func (h *Handlers) LoginHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("internal/templates/login.html"))
	return tmpl.Execute(c.Response().Writer, nil)

}
