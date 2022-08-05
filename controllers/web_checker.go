package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/parmcoder/website-checker-backend/configs"
)

//go:generate mockgen -destination=mocks/web_checker.go -package=mocks

func (s *ServerImpl) CheckHealthCsv(ctx echo.Context) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	lines, err := (*s.checker).ExtractLinesFromCsv(file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	var listOfSites configs.WebsiteList

	listOfSites.Rows = lines

	ups, down, duration, err := (*s.checker).PerformCheck(&listOfSites.Rows)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := configs.WebsiteListResponse{
		Ups:      ups,
		Downs:    down,
		Duration: duration,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (s *ServerImpl) CheckHealth(ctx echo.Context) error {

	var listOfSites configs.WebsiteList

	err := ctx.Bind(&listOfSites)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	ups, down, duration, err := (*s.checker).PerformCheck(&listOfSites.Rows)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := configs.WebsiteListResponse{
		Ups:      ups,
		Downs:    down,
		Duration: duration,
	}

	return ctx.JSON(http.StatusOK, response)
}
