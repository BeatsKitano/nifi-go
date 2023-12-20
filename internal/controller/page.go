package controller

import (
	"nifi-go/internal/service"

	"github.com/labstack/echo/v4"
)

type PageController struct {
	query *service.PageQuery
}

func NewPageController(query *service.PageQuery) *PageController {
	c := &PageController{
		query: query,
	}

	return c
}

func (c *PageController) RegisterRouter(router *echo.Echo) {
	router.GET("/pages", c.list)
}

func (c *PageController) list(e echo.Context) error {
	str, err := c.query.List()
	if err != nil {
		return err
	}
	return e.JSON(200, str)
}
