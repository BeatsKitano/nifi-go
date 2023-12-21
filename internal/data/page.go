package data

import (
	"context"
	"nifi-go/internal/app/nifigo/service"
)

var _ service.PageRepo = (*pageRepo)(nil)

type pageRepo struct {
	data *Data
}

func NewPageRepo(data *Data) service.PageRepo {
	return &pageRepo{
		data: data,
	}
}

func (r *pageRepo) List(ctx context.Context) (int32, []*service.Page, error) {
	return 4, []*service.Page{
		{
			ID:   1,
			Name: "name01",
			Age:  18,
		},
		{
			ID:   2,
			Name: "name02",
			Age:  19,
		},
		{
			ID:   3,
			Name: "name03",
			Age:  20,
		},
		{
			ID:   4,
			Name: "name04",
			Age:  21,
		},
	}, nil

}
