//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"

	"github.com/parmcoder/website-checker-backend/controllers"
	"github.com/parmcoder/website-checker-backend/repositories"
	"github.com/parmcoder/website-checker-backend/services"
)

func initializeCheckerService() services.CheckerService {
	wire.Build(repositories.NewWebCheckerRepository, services.NewCheckerService)
	return services.CheckerServiceImpl{}
}

func InitializeServerParams() controllers.ServerImplParams {
	wire.Build(wire.Struct(new(controllers.ServerImplParams), "*"), initializeCheckerService)
	return controllers.ServerImplParams{}
}

func InitializeServer() controllers.Server {
	wire.Build(controllers.NewServer, InitializeServerParams)
	return &controllers.ServerImpl{}
}
