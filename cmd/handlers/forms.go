package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleClearForm(c echo.Context) error {
	return c.Render(http.StatusOK, c.QueryParam("template"), newFormData())
}
