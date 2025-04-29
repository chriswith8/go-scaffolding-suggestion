package validator

import (
	"sync"

	"github.com/go-playground/validator"
)

var (
	once      sync.Once
	singleton *validator.Validate
)

func Validator() *validator.Validate {
	once.Do(func() {
		singleton = validator.New()
	})
	return singleton
}
