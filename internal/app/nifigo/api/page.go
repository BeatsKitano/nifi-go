package api

import (
	"fmt"
	"nifi-go/internal/app/nifigo/service"

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
	router.POST("/tokens", c.createToken)
}

func (c *PageController) createToken(e echo.Context) error {
	type Params struct {
		AppKey           string `json:"appKey"`
		RequestTimestamp string `json:"requestTimestamp"`
		Sign             string `json:"sign"`
	}

	var params Params
	if err := e.Bind(&params); err != nil {
		fmt.Println(err, "=====")
		return err
	}

	fmt.Println(params)

	return e.JSON(200, map[string]interface{}{
		"token":   "token203920020299202020200202",
		"success": true,
	})
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /pages [get]
func (c *PageController) list(e echo.Context) error {

	type Params struct {
		PageSize    int32 `query:"page_size"`
		PageNum     int32 `query:"page_num"`
		HasNextPage bool  `query:"hasNextPage"`
	}

	var params Params
	if err := e.Bind(&params); err != nil {
		return err
	}

	many, err := c.query.List(e.Request().Context())
	if err != nil {
		return err
	}

	if params.PageNum > 1 {
		// 去掉第一个元素
		many = many[1:]

		for _, v := range many {
			v.ID = v.ID + 10
		}
	}

	return e.JSON(200, map[string]interface{}{
		"data": many,
	})
}
