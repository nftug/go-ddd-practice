package repository

import "go-ddd-practice/domain/vo"

type Repository[T any] interface {
	GetById(id *vo.Id) (*T, error)
	Insert(entity *T) (*vo.Id, error)
}
