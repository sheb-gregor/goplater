// generated by goplater -type=WeekDay; DO NOT EDIT
package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func init() {
	// stub usage of json for situation when
	// (Un)MarshalJSON methods will be omitted
	_ = json.Delim('s')

	_ = driver.Bool
}

var ErrWeekDayInvalid = errors.New("WeekDay is invalid")

func init() {
	var v WeekDay
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_WeekDayNameToValue = map[string]WeekDay{
			interface{}(Monday).(fmt.Stringer).String():    Monday,
			interface{}(Tuesday).(fmt.Stringer).String():   Tuesday,
			interface{}(Wednesday).(fmt.Stringer).String(): Wednesday,
			interface{}(Thursday).(fmt.Stringer).String():  Thursday,
			interface{}(Friday).(fmt.Stringer).String():    Friday,
			interface{}(Saturday).(fmt.Stringer).String():  Saturday,
			interface{}(Sunday).(fmt.Stringer).String():    Sunday,
		}
	}
}

var _WeekDayNameToValue = map[string]WeekDay{
	"Monday":    Monday,
	"Tuesday":   Tuesday,
	"Wednesday": Wednesday,
	"Thursday":  Thursday,
	"Friday":    Friday,
	"Saturday":  Saturday,
	"Sunday":    Sunday,
}

var _WeekDayValueToName = map[WeekDay]string{
	Monday:    "Monday",
	Tuesday:   "Tuesday",
	Wednesday: "Wednesday",
	Thursday:  "Thursday",
	Friday:    "Friday",
	Saturday:  "Saturday",
	Sunday:    "Sunday",
}

// Validate verifies that value is predefined for WeekDay.
func (r WeekDay) Validate() error {
	_, ok := _WeekDayValueToName[r]
	if !ok {
		return ErrWeekDayInvalid
	}
	return nil
}

// MarshalJSON is generated so WeekDay satisfies json.Marshaler.
func (r WeekDay) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _WeekDayValueToName[r]
	if !ok {
		return nil, fmt.Errorf("WeekDay(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so WeekDay satisfies json.Unmarshaler.
func (r *WeekDay) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WeekDay should be a string, got %s", string(data))
	}
	v, ok := _WeekDayNameToValue[s]
	if !ok {
		return fmt.Errorf("WeekDay(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so WeekDay satisfies db row driver.Valuer.
func (r WeekDay) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so WeekDay satisfies db row driver.Scanner.
func (r *WeekDay) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i WeekDay
	err := json.Unmarshal(source, &i)
	if err != nil {
		return errors.New("WeekDay: can't unmarshal column data")
	}

	*r = i
	return nil
}