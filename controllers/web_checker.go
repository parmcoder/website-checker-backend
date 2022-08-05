package controllers

import (
	"encoding/csv"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/parmcoder/website-checker-backend/configs"
)

func (s *ServerImpl) CheckHealthCsv(ctx echo.Context) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	defer src.Close()

	csvReader := csv.NewReader(src)
	records, err := csvReader.ReadAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	lines, err := (*s.checker).ExtractLinesFromCsv(records)
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
