# goplater

**goplater** is a tool for automating the creation of methods that satisfy some built-in interfaces or those defined in the templates.

# Usage 

**goplater** can be run from console in directory with target type or with help `go generate`:

```bash
$ goplater -type=MyType
``` 

Or

``` go
package main

//go:generate goplater -type=MyType
type MyType int 

const (
	MyTypeA MyType   = iota
	MyTypeB
	MyTypeC
)
```

List of arguments:

| Flag | Values | Description |
| ---- | ------ | ----------- |
| type | string | The name of the target type or types for code generation |
| transform | snake, kebab, space, none | A rule describing the strategy for converting constant names to a string. Default: none|
| tprefix | true, false | add type name prefix into string values or not. Default: false |
| prefix | string |  A prefix to be added to the output file |
| suffix | string |  A suffix to be added to the output. Default: "_enums"|

# ToDo

- [ ] Improve documentation, add tests
- [ ] Add string types support 
- [ ] Add bitmap types support
- [ ] Add custom template support

##### OLD
___

Given the name of a (signed or unsigned) integer type T that has constants
defined, jsonenums will create a new self-contained Go source file implementing

```
func (t T) MarshalJSON() ([]byte, error)
func (t *T) UnmarshalJSON([]byte) error
```

The file is created in the same package and directory as the package that
defines T. It has helpful defaults designed for use with go generate.

jsonenums is a simple implementation of a concept and the code might not be the
most performant or beautiful to read.

For example, given this snippet,

```Go
package painkiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
```

running this command

```
jsonenums -type=Pill
```

in the same directory will create the file `pill_jsonenums.go`, in package
`painkiller`, containing a definition of

```
func (r Pill) MarshalJSON() ([]byte, error)
func (r *Pill) UnmarshalJSON([]byte) error
```

`MarshalJSON` will translate the value of a `Pill` constant to the `[]byte`
representation of the respective constant name, so that the call
`json.Marshal(painkiller.Aspirin)` will return the bytes `[]byte("\"Aspirin\"")`.

`UnmarshalJSON` performs the opposite operation; given the `[]byte`
representation of a `Pill` constant it will change the receiver to equal the
corresponding constant. So given `[]byte("\"Aspirin\"")` the receiver will
change to `Aspirin` and the returned error will be `nil`.

Typically this process would be run using go generate, like this:

```
//go:generate jsonenums -type=Pill
```

If multiple constants have the same value, the lexically first matching name
will be used (in the example, Acetaminophen will print as "Paracetamol").

With no arguments, it processes the package in the current directory. Otherwise,
the arguments must name a single directory holding a Go package or a set of Go
source files that represent a single Go package.

The `-type` flag accepts a comma-separated list of types so a single run can
generate methods for multiple types. The default output file is t_jsonenums.go,
where t is the lower-cased name of the first type listed. The suffix can be
overridden with the `-suffix` flag and a prefix may be added with the `-prefix` 
flag.

This is not an official Google product (experimental or otherwise), it is just code that happens to be owned by Google.

# License

This package is based on https://github.com/campoy/jsonenums.