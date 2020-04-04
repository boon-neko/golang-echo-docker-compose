package test

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" `
	Email string `json:"email"`
}

func RoutingCheck() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, u)
	}

}
