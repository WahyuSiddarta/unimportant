package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a API) TestLogin(c echo.Context) error {
	fmt.Println("test hit")
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login Successful",
	})
}
