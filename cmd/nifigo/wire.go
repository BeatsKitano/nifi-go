//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"nifi-go/internal/app/nifigo/api"
	"nifi-go/internal/app/nifigo/server"
	"nifi-go/internal/app/nifigo/service"
	"nifi-go/internal/data"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// wireApp init
func wireApp() *server.HttpServer {
	panic(wire.Build(api.ApiSet, service.ServiceSet, server.ServerSet, data.DataSet, echo.New))
}
