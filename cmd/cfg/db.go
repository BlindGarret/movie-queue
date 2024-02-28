package cfg

import (
	"github.com/BlindGarret/movie-queue/ent"
	"github.com/labstack/echo/v4"
)

const (
	DBXContextKey = "__dbx__"
)

func DBXMiddleware(client *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DBXContextKey, client)
			return next(c)
		}
	}
}
