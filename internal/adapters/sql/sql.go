package sql

import (
	"gorm.io/gorm"
)

type (
	Adapter[T interface{}] interface {
		Create(e T) *gorm.DB
	}
	sqlClientAdapter[T interface{}] struct {
		conn *gorm.DB
	}
)

func NewSQLAdapter[T interface{}](conn *gorm.DB) Adapter[T] {
	return &sqlClientAdapter[T]{conn: conn}
}

func (s sqlClientAdapter[T]) Create(e T) *gorm.DB {
	return s.conn.Create(e)
}
