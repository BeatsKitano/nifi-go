package service

import "context"

type PageQuery struct {
	page PageRepo
}

func NewPageQuery(page PageRepo) *PageQuery {
	return &PageQuery{
		page: page,
	}
}

func (q *PageQuery) List(ctx context.Context) ([]*Page, error) {
	return q.page.List(ctx)
}
