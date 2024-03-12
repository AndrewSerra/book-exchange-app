package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// The FullDate type is used for the format
// YYYY-MM-DD and will UnMarshall properly
type FullDate struct {
	time.Time
}

func (d *FullDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *FullDate) MarshalJSON() ([]byte, error) {
	dateStr := fmt.Sprintf("\"%s\"", d.Time.Format("2006-01-02"))
	return []byte(dateStr), nil
}

func (d *FullDate) Scan(value interface{}) error {
	switch value.(type) {
	case time.Time:
		(*d).Time = value.(time.Time)
	case []uint8:
		newTime, err := time.Parse("2006-01-02", string(value.([]byte)))
		(*d).Time = newTime
		if err != nil {
			return &UnknownError{Err: err}
		}
	default:
		return &UnrecognizedTypeError{
			Data: value,
			Err:  errors.New("unrecognized type"),
		}
	}
	return nil
}

func (d FullDate) Value() (driver.Value, error) {
	return d.Time.Format("2006-01-02"), nil
}

// The MonthYearDate type is used for the format
// YYYY-MM-DD and will UnMarshall properly
type MonthYearDate struct {
	time.Time
}

func (d *MonthYearDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *MonthYearDate) MarshalJSON() ([]byte, error) {
	dateStr := fmt.Sprintf("\"%s\"", d.Time.Format("2006-01"))
	return []byte(dateStr), nil
}

func (d *MonthYearDate) Scan(value interface{}) error {
	switch value.(type) {
	case time.Time:
		(*d).Time = value.(time.Time)
	case []uint8:
		newTime, err := time.Parse("2006-01", string(value.([]byte)))
		(*d).Time = newTime
		if err != nil {
			return &UnknownError{Err: err}
		}
	default:
		return &UnrecognizedTypeError{
			Data: value,
			Err:  errors.New("unrecognized type"),
		}
	}
	return nil
}

func (d MonthYearDate) Value() (driver.Value, error) {
	return d.Time.Format("2006-01"), nil
}