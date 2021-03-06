package db

import (
	"database/sql/driver"
	"errors"
	"strings"
)

// List is a custom sql/driver type to handle list columns
type List []string

// Scan - Implement the database/sql/driver Scanner interface
func (a *List) Scan(value interface{}) error {
	if value == nil {
		*a = List([]string{})
		return nil
	}
	if bv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := bv.(string); ok {
			if len(v) == 0 {
				*a = List([]string{})
			}
			if len(v) > 0 {
				*a = List(strings.Split(v, ","))
			}
			return nil
		}
	}
	return errors.New("failed to scan List")
}

// Value - Implement the database/sql/driver Valuer interface
func (a List) Value() (driver.Value, error) {
	return strings.Join(a, ","), nil
}

func (a *List) String() string {
	v, _ := a.Value()
	return v.(string)
}
