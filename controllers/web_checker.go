package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *ServerImpl) CheckHealth(ctx echo.Context) error {
	var listOfSites []string

	err := (*s.checker).PerformCheck(&listOfSites)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, "Good")
}
