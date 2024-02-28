package handlers

import (
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	// Register all handlers here

	// todo: sorting (mutable)

	// Index
	e.GET("/", handleIndex)
	e.GET("/shows", handleGetShows)
	e.POST("/shows", handleAddShow)
	e.DELETE("/shows/:id", handleDeleteShow)

	// Forms
	e.GET("/forms/clear", handleClearForm)
}
