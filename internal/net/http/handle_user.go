package http

import (
	"net/http"
	"os/user"

	"github.com/labstack/echo/v4"
)

func (h Handler) handleEnrol() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		user := user.User{}
		_ = c.Bind(&user)
		newUser, err := h.users.Enrol(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, newUser)
	}
}

func (h Handler) handleStudentInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		user := user.User{}
		_ = c.Bind(&user)
		newUser, err := h.users.Enrol(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, newUser)
	}
}
