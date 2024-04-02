package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type BitBool bool

func (b *BitBool) Scan(value interface{}) error {
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}

func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

type Genre []string

func (g *Genre) Scan(value interface{}) error {
	content := string(value.([]byte))
	*g = strings.Split(content, ",")
	return nil
}

func (g Genre) Value() (driver.Value, error) {
	return strings.Join(g, ","), nil
}

type YearDate int

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
