package service

import "context"

type Page struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

type PageRepo interface {
	List(ctx context.Context) (int32, []*Page, error)
}
