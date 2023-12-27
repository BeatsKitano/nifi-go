package data

import (
	"context"
	"nifi-go/internal/app/nifigo/service"
	"time"
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

func (r *pageRepo) List(ctx context.Context) ([]*service.Page, error) {
	t := time.Now().Format("2006-01-02 15:04:05")
	return []*service.Page{
		{
			ID:   1,
			Name: "name0111",
			Age:  18,
			Time: t,
		},
		{
			ID:   2,
			Name: "name0222",
			Age:  19,
			Time: t,
		},
		{
			ID:   3,
			Name: "name0333",
			Age:  20,
			Time: t,
		},
		{
			ID:   4,
			Name: "name0444",
			Age:  21,
			Time: t,
		},
	}, nil

}
