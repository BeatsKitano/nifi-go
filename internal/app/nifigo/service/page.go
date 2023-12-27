package service

import "context"

type Page struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
	Time string `json:"time"`
}

type PageRepo interface {
	List(ctx context.Context) ([]*Page, error)
}
