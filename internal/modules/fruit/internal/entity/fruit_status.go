package entity

import (
	"database/sql/driver"

	"github.com/chriswith8/go-scaffolding-suggestion/internal/utils/constants"
	"github.com/chriswith8/go-scaffolding-suggestion/pkg/common"
)

type FruitStatus int

const (
	Undefined FruitStatus = iota
	Ripe
	Unripe
	Rotten
	Destroyed
)

var fruitStatusMap = map[string]FruitStatus{
	"Ripe":                    Ripe,
	"Unripe":                  Unripe,
	"Rotten":                  Rotten,
	"Destroyed":               Destroyed,
	constants.UndefinedStatus: Undefined,
	constants.Blank:           Undefined,
}

func ToFruitStatus(s string) (FruitStatus, error) {
	return common.ToEnum(s, fruitStatusMap)
}

func (e *FruitStatus) UnmarshalJSON(bytes []byte) (err error) {
	return common.UnmarshalJSON(e, ToFruitStatus, bytes, true)
}

func (e FruitStatus) MarshalJSON() ([]byte, error) {
	return common.MarshalJSON(e)
}

func (e FruitStatus) Value() (driver.Value, error) {
	return e.String(), nil
}

func (e *FruitStatus) Scan(src any) (err error) {
	return common.Scan(e, ToFruitStatus, src, true)
}

func (e FruitStatus) String() string {
	switch e {
	case Ripe:
		return "Ripe"
	case Unripe:
		return "Unripe"
	case Rotten:
		return "Rotten"
	case Destroyed:
		return "Destroyed"
	default:
		return constants.UndefinedStatus
	}
}
