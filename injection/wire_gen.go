// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/parmcoder/website-checker-backend/controllers"
	"github.com/parmcoder/website-checker-backend/services"
)

// Injectors from wire.go:

func initializeCheckerService() services.CheckerService {
	checkerService := services.NewCheckerService()
	return checkerService
}

func InitializeServerParams() controllers.ServerImplParams {
	checkerService := initializeCheckerService()
	serverImplParams := controllers.ServerImplParams{
		checker: checkerService,
	}
	return serverImplParams
}

func initializeServer() controllers.Server {
	serverImplParams := InitializeServerParams()
	server := controllers.NewServer(serverImplParams)
	return server
}