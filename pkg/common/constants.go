package common

import (
	"errors"
	"fmt"
)

type (
	Stringer interface {
		String() string
	}
	mapperFn[T any] func(string) (T, error)
)

func ToEnum[T any](s string, mapper map[string]T) (T, error) {
	if option, ok := mapper[s]; ok {
		return option, nil
	}
	var zero T
	return zero, errors.New("invalid value enum value: " + s)
}

func MarshalJSON(c Stringer) ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

func UnmarshalJSON[T any](c *T, mapper mapperFn[T], b []byte, failOnMapperError bool) (err error) {
	if len(b) < 2 {
		return errors.New("unexpected end of JSON input")
	}
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	if *c, err = mapper(string(b)); err != nil && failOnMapperError {
		return err
	}

	return nil
}

func Scan[T any](t *T, mapper mapperFn[T], src any, nilIsValid bool) (err error) {
	if src == nil {
		if nilIsValid {
			return
		} else {
			return fmt.Errorf("nil is not a valid value")
		}
	}

	switch typedSrc := src.(type) {
	case []byte:
		*t, err = mapper(string(typedSrc))
	case string:
		*t, err = mapper(typedSrc)
	default:
		err = fmt.Errorf("value is not a string or []byte: %v", src)
	}

	return
}
