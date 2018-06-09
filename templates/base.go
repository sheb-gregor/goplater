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

package templates

import "text/template"

type CodeTemplate struct {
	Name   string
	Raw    string
	Parsed *template.Template
}

var Base = []CodeTemplate{
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

func (r *CodeTemplate) parse() {
	r.Parsed = template.Must(template.New(r.Name).Parse(r.Raw))
}

func init() {
	for i := range Base {
		Base[i].parse()
	}
}

var (
	baseTmplRaw = `
// generated by goplater {{.Command}}; DO NOT EDIT
package {{.PackageName}}

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

{{range $typename, $values := .TypesAndValues}}
var Err{{$typename}}Invalid = errors.New("{{$typename}} is invalid")
{{end}}
`

	nameToValueTmplRaw = `
{{range $typename, $values := .TypesAndValues}}
var def{{$typename}}NameToValue = map[string]{{$typename}} {
        {{range $values}}"{{.Str}}": {{.Name}},
        {{end}}
    }

{{end}}
`

	valueToNameTmplRaw = `
{{range $typename, $values := .TypesAndValues}}
var def{{$typename}}ValueToName = map[{{$typename}}]string {
        {{range $values}}{{.Name}}: "{{.Str}}",
        {{end}}
    }

{{end}}
`
	stringTmplRaw = `
{{range $typename, $values := .TypesAndValues}}
// String is generated so {{$typename}} satisfies fmt.Stringer.
func (r {{$typename}}) String() string {
    s, ok := def{{$typename}}ValueToName[r]
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
    _, ok := def{{$typename}}ValueToName[r]
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
    s, ok := def{{$typename}}ValueToName[r]
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
        return fmt.Errorf("{{$typename}}: should be a string, got %s", string(data))
    }
    v, ok := def{{$typename}}NameToValue[s]
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
    switch v := src.(type) {
    case string:
        val, ok := def{{$typename}}NameToValue[v]
        if !ok {
            return errors.New("{{$typename}}: can't unmarshal column data")
        }
        *r = val
        return nil
    case []byte:
        var i {{$typename}}
        err := json.Unmarshal(v, &i)
        if err != nil {
            return errors.New("{{$typename}}: can't unmarshal column data")
        }
    
        *r = i
        return nil
    case int, int8, int32, int64, uint, uint8, uint32, uint64:
        ni := sql.NullInt64{}
        err := ni.Scan(v)
        if err != nil {
            return errors.New("{{$typename}}: can't scan column data into int64")
        }
    
        *r = {{$typename}}(ni.Int64)
        return nil
    }
    return errors.New("{{$typename}}: invalid type")
}
{{end}}
`
)