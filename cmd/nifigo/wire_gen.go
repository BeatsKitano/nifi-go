// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/labstack/echo/v4"
	"nifi-go/internal/controller"
	"nifi-go/internal/server"
	"nifi-go/internal/service"
)

// Injectors from wire.go:

// wireApp init
func wireApp() *server.HttpServer {
	echoEcho := echo.New()
	pageQuery := service.NewPageQuery()
	pageController := controller.NewPageController(pageQuery)
	httpServer := server.NewHTTPServer(echoEcho, pageController)
	return httpServer
}