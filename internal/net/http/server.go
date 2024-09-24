package http

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.tmp/echo/internal/notify"
	"go.tmp/echo/internal/user"
)

type Server interface {
	Start(address string) error
}

func NewServer(connSQL *sqlx.DB) Server {
	h := Handler{
		users:         &user.DB{RWC: connSQL},
		notifications: &notify.DB{RWC: connSQL},
	}
	e := echo.New()
	e.POST("/users", h.handleEnrol())
	e.POST("/student/:id", h.handleStudentInfo())
	e.GET("/ready", h.handleReady( /* set route deps */ ))
	return e
}

type Handler struct {
	users         *user.DB
	notifications *notify.DB
}

func (h Handler) handleReady() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "READY\n")
	}
}
