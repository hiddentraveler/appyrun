package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, h.server.DB.Health())
}
