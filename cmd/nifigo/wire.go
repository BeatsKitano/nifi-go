//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"nifi-go/internal/controller"
	"nifi-go/internal/server"
	"nifi-go/internal/service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// wireApp init
func wireApp() *server.HttpServer {
	panic(wire.Build(controller.ControllerSet, service.ServiceSet, server.ServerSet, echo.New))
}
