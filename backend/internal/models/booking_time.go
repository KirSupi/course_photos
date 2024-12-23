package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const Format = "2006-01-02 15"

type BookingTime time.Time

func (d *BookingTime) Unix() int64 {
	return time.Time(*d).Truncate(time.Hour).Unix()
}

func (d *BookingTime) String() string {
	return time.Time(*d).Truncate(time.Hour).Format(Format)
}

func (d *BookingTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
func (d *BookingTime) UnmarshalJSON(data []byte) error {
	dateStr := ""

	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}

	return d.UnmarshalText([]byte(dateStr))
}

func (d *BookingTime) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}
func (d *BookingTime) UnmarshalText(text []byte) error {
	timeTime, err := time.Parse(Format, string(text))
	if err != nil {
		return err
	}

	*d = BookingTime(timeTime.Truncate(time.Hour))

	return nil
}

func (d *BookingTime) Scan(value any) error {
	if timeTime, ok := value.(time.Time); ok {
		*d = BookingTime(timeTime.Truncate(time.Hour))

		return nil
	} else {
		return errors.New(fmt.Sprintf("unsupported type for (d *Date) Scan()"))
	}
}
func (d *BookingTime) Value() (val driver.Value, err error) {
	return d.MarshalText()
}
