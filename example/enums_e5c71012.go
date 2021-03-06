// generated by goplater enum --type ShirtSize,WeekDay --merge true; DO NOT EDIT
package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func init() {
	// stub usage of json for situation when
	// (Un)MarshalJSON methods will be omitted
	_ = json.Delim('s')

	// stub usage of sql/driver for situation when
	// Scan/Value methods will be omitted
	_ = driver.Bool
	_ = sql.LevelDefault
}

var ErrShirtSizeInvalid = errors.New("ShirtSize is invalid")

var defShirtSizeNameToValue = map[string]ShirtSize{
	"NA": NA,
	"XS": XS,
	"S":  S,
	"M":  M,
	"L":  L,
	"XL": XL,
}

// String is generated so ShirtSize satisfies fmt.Stringer.
func (r ShirtSize) String() string {
	s, ok := defShirtSizeValueToName[r]
	if !ok {
		return fmt.Sprintf("ShirtSize(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for ShirtSize.
func (r ShirtSize) Validate() error {
	_, ok := defShirtSizeValueToName[r]
	if !ok {
		return ErrShirtSizeInvalid
	}
	return nil
}

// MarshalJSON is generated so ShirtSize satisfies json.Marshaler.
func (r ShirtSize) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defShirtSizeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("ShirtSize(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ShirtSize satisfies json.Unmarshaler.
func (r *ShirtSize) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ShirtSize: should be a string, got %s", string(data))
	}
	v, ok := defShirtSizeNameToValue[s]
	if !ok {
		return fmt.Errorf("ShirtSize(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so ShirtSize satisfies db row driver.Valuer.
func (r ShirtSize) Value() (driver.Value, error) {
	s, ok := defShirtSizeValueToName[r]
	if !ok {
		return nil, nil
	}
	return s, nil
}

// Value is generated so ShirtSize satisfies db row driver.Scanner.
func (r *ShirtSize) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		val, _ := defShirtSizeNameToValue[v]
		*r = val
		return nil
	case []byte:
		var i ShirtSize
		err := json.Unmarshal(v, &i)
		if err != nil {
			return errors.New("ShirtSize: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(v)
		if err != nil {
			return errors.New("ShirtSize: can't scan column data into int64")
		}

		*r = ShirtSize(ni.Int64)
		return nil
	}
	return errors.New("ShirtSize: invalid type")
}

var ErrWeekDayInvalid = errors.New("WeekDay is invalid")

var defWeekDayNameToValue = map[string]WeekDay{
	"Monday":    Monday,
	"Tuesday":   Tuesday,
	"Wednesday": Wednesday,
	"Thursday":  Thursday,
	"Friday":    Friday,
	"Saturday":  Saturday,
	"Sunday":    Sunday,
}

// String is generated so WeekDay satisfies fmt.Stringer.
func (r WeekDay) String() string {
	s, ok := defWeekDayValueToName[r]
	if !ok {
		return fmt.Sprintf("WeekDay(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for WeekDay.
func (r WeekDay) Validate() error {
	_, ok := defWeekDayValueToName[r]
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
	s, ok := defWeekDayValueToName[r]
	if !ok {
		return nil, fmt.Errorf("WeekDay(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so WeekDay satisfies json.Unmarshaler.
func (r *WeekDay) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WeekDay: should be a string, got %s", string(data))
	}
	v, ok := defWeekDayNameToValue[s]
	if !ok {
		return fmt.Errorf("WeekDay(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so WeekDay satisfies db row driver.Valuer.
func (r WeekDay) Value() (driver.Value, error) {
	s, ok := defWeekDayValueToName[r]
	if !ok {
		return nil, nil
	}
	return s, nil
}

// Value is generated so WeekDay satisfies db row driver.Scanner.
func (r *WeekDay) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		val, _ := defWeekDayNameToValue[v]
		*r = val
		return nil
	case []byte:
		var i WeekDay
		err := json.Unmarshal(v, &i)
		if err != nil {
			return errors.New("WeekDay: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(v)
		if err != nil {
			return errors.New("WeekDay: can't scan column data into int64")
		}

		*r = WeekDay(ni.Int64)
		return nil
	}
	return errors.New("WeekDay: invalid type")
}
