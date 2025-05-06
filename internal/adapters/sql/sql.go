package sql

import (
	"gorm.io/gorm"
)

type (
	Adapter[T interface{}] interface {
		Create(e T) *gorm.DB
		Save(e T) *gorm.DB
		First(e *T) *gorm.DB
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

func (s sqlClientAdapter[T]) Save(e T) *gorm.DB {
	return s.conn.Save(e)
}

func (s sqlClientAdapter[T]) First(e *T) *gorm.DB {
	return s.conn.First(e)
}
