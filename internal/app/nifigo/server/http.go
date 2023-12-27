package server

import (
	"nifi-go/internal/app/nifigo/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	e *echo.Echo
}

// 必须持有/依赖 controller 否则没有引用，wire 会认为该 provider 无用
func NewHTTPServer(e *echo.Echo, page *api.PageController) *HttpServer {

	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		CustomTimeFormat: "2006-01-02 15:04:05",
		Format:           "time=${time_custom}, method=${method}, status=${status}, uri=${uri}\n",
	}))

	page.RegisterRouter(e)

	return &HttpServer{
		e: e,
	}
}

func (s *HttpServer) Start(port string) error {
	return s.e.Start(port)
}
