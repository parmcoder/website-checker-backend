package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/parmcoder/website-checker-backend/configs"
)

func (s *ServerImpl) CheckHealth(ctx echo.Context) error {
	var listOfSites configs.WebsiteList

	err := ctx.Bind(&listOfSites)
	if err != nil {
		return err
	}

	ups, down, duration, err := (*s.checker).PerformCheck(&listOfSites.Rows)
	if err != nil {
		return err
	}

	response := configs.WebsiteListResponse{
		Ups:      ups,
		Downs:    down,
		Duration: duration,
	}

	return ctx.JSON(http.StatusOK, response)
}
