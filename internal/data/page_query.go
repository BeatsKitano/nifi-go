package data

import (
	"nifi-go/internal/app/nifigo/service"
)

var _ service.PageQueryRepo = (*pageQueryRepo)(nil)

type pageQueryRepo struct {
	data *Data
}

func NewPageQueryRepo(data *Data) service.PageQueryRepo {
	return &pageQueryRepo{
		data: data,
	}
}
