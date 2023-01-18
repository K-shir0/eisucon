package echo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	db *sqlx.DB
}

var indent = "	"

func JSONMessage(ctx echo.Context, code int, msg string) error {
	return ctx.JSONPretty(code, map[string]string{"message": msg}, indent)
}

func JSONPretty(ctx echo.Context, code int, i interface{}) error {
	return ctx.JSONPretty(code, i, indent)
}
