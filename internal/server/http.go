package server

import (
	"nifi-go/internal/controller"

	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	e *echo.Echo
}

// 必须持有/依赖 controller 否则没有引用，wire 会认为该 provider 无用
func NewHTTPServer(e *echo.Echo, page *controller.PageController) *HttpServer {
	page.RegisterRouter(e)

	return &HttpServer{
		e: e,
	}
}

func (s *HttpServer) Start(port string) error {
	return s.e.Start(port)
}
