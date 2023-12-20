package service

type PageQueryRepo interface {
}

type PageQuery struct {
	query PageQueryRepo
}

func NewPageQuery(query PageQueryRepo) *PageQuery {
	return &PageQuery{
		query: query,
	}
}

func (q *PageQuery) List() (string, error) {
	return "page list", nil
}
