package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) HomeHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
