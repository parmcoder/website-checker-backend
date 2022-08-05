package controllers

import (
	echo "github.com/labstack/echo/v4"

	"github.com/parmcoder/website-checker-backend/services"
)

type Server interface {
	CheckHealth(echo.Context) error
}

type ServerImplParams struct {
	Checker services.CheckerService
}

type ServerImpl struct {
	checker *services.CheckerService
}

func NewServer(params ServerImplParams) Server {
	return &ServerImpl{
		checker: &params.Checker,
	}
}
