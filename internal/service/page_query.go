package service

type PageQuery struct {
}

func NewPageQuery() *PageQuery {
	return &PageQuery{}
}

func (q *PageQuery) List() (string, error) {
	return "page list", nil
}
