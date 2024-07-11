package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Update User
// (PUT /api/v1/users/{userId})
func (s *Controller) GetHello(ctx echo.Context) error {
	return returnNotImplemented(ctx)
}

func returnNotImplemented(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, "Not Implemented")
}
