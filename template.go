// Copyright 2017 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

// Added as a .go file to avoid embedding issues of the template.

package main

import "text/template"

type TemplatePart struct {
	Name   string
	Raw    string
	Parsed *template.Template
}

var templateParts = []TemplatePart{
	{Name: "base", Raw: baseTmplRaw},
	{Name: "NameToValue", Raw: nameToValueTmplRaw},
	{Name: "ValueToName", Raw: valueToNameTmplRaw},
	{Name: "String", Raw: stringTmplRaw},
	{Name: "Validate", Raw: validateTmplRaw},
	{Name: "MarshalJSON", Raw: marshalJSONTmplRaw},
	{Name: "UnmarshalJSON", Raw: unmarshalJSONTmplRaw},
	{Name: "Value", Raw: rowValueTmplRaw},
	{Name: "Scan", Raw: rowScanTmplRaw},
}

func (r *TemplatePart) parse() {
	r.Parsed = template.Must(template.New(r.Name).Parse(r.Raw))
}

func init() {
	for i := range templateParts {
		templateParts[i].parse()
	}
}

var (
	baseTmplRaw = `
// generated by jsonenums {{.Command}}; DO NOT EDIT
package {{.PackageName}}

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

{{range $typename, $values := .TypesAndValues}}
var Err{{$typename}}Invalid = errors.New("{{$typename}} is invalid")

func init() {
    var v {{$typename}}
    if _, ok := interface{}(v).(fmt.Stringer); ok {
        _{{$typename}}NameToValue = map[string]{{$typename}} {
            {{range $values}}interface{}({{.Name}}).(fmt.Stringer).String(): {{.Name}},
            {{end}}
        }
    }
}
{{end}}

`

	nameToValueTmplRaw = `
{{range $typename, $values := .TypesAndValues}}
var _{{$typename}}NameToValue = map[string]{{$typename}} {
        {{range $values}}"{{.Str}}": {{.Name}},
        {{end}}
    }

{{end}}
`

	valueToNameTmplRaw = `
{{range $typename, $values := .TypesAndValues}}
var _{{$typename}}ValueToName = map[{{$typename}}]string {
        {{range $values}}{{.Name}}: "{{.Str}}",
        {{end}}
    }

{{end}}
`
	stringTmplRaw = `
{{range $typename, $values := .TypesAndValues}}
// String is generated so {{$typename}} satisfies fmt.Stringer.
func (r {{$typename}}) String() string {
    s, ok := _{{$typename}}ValueToName[r]
    if !ok {
        return fmt.Sprintf("{{$typename}}(%d)", r)
    }
    return s
}
{{end}}

`

	validateTmplRaw = `
{{range $typename, $values := .TypesAndValues}}
// Validate verifies that value is predefined for {{$typename}}.
func (r {{$typename}}) Validate() error {
    _, ok := _{{$typename}}ValueToName[r]
    if !ok {
        return Err{{$typename}}Invalid
    }
    return nil
}
{{end}}

`

	marshalJSONTmplRaw = `
{{range $typename, $values := .TypesAndValues}}

// MarshalJSON is generated so {{$typename}} satisfies json.Marshaler.
func (r {{$typename}}) MarshalJSON() ([]byte, error) {
    if s, ok := interface{}(r).(fmt.Stringer); ok {
        return json.Marshal(s.String())
    }
    s, ok := _{{$typename}}ValueToName[r]
    if !ok {
        return nil, fmt.Errorf("{{$typename}}(%d) is invalid value", r)
    }
    return json.Marshal(s)
}

{{end}}
`

	unmarshalJSONTmplRaw = `
{{range $typename, $values := .TypesAndValues}}

// UnmarshalJSON is generated so {{$typename}} satisfies json.Unmarshaler.
func (r *{{$typename}}) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return fmt.Errorf("{{$typename}} should be a string, got %s", string(data))
    }
    v, ok := _{{$typename}}NameToValue[s]
    if !ok {
        return fmt.Errorf("{{$typename}}(%q) is invalid value", s)
    }
    *r = v
    return nil
}

{{end}}
`

	rowValueTmplRaw = `
{{range $typename, $values := .TypesAndValues}}

// Value is generated so {{$typename}} satisfies db row driver.Valuer.
func (r {{$typename}}) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}
{{end}}
`
	rowScanTmplRaw = `
{{range $typename, $values := .TypesAndValues}}

// Value is generated so {{$typename}} satisfies db row driver.Scanner.
func (r *{{$typename}}) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i {{$typename}}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return errors.New("{{$typename}}: can't unmarshal column data")
	}

	*r = i
	return nil
}
{{end}}
`
)
