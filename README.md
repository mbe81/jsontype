# jsontype for Go

[![GoDoc](https://godoc.org/github.com/mbe81/jsontype?status.svg)](https://godoc.org/github.com/mbe81/jsontype)
[![Go Report Card](https://goreportcard.com/badge/github.com/mbe81/jsontype)](https://goreportcard.com/report/github.com/mbe81/jsontype)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

The `jsontype` package provides some basic types for handling JSON unmarshaling in scenarios where fields may be null or absent, such as when processing data from PATCH requests. The types are inspired by the `sql.Null*` types from the `database/sql` package.

## Overview

The types provided by this package are designed to work seamlessly with the `encoding/json` package from the Go standard library. They implement the `json.Unmarshaler` and `json.Marshaler` interfaces, making it easy to manage fields that can either be present or absent, and may or may not contain null values.

Each type includes two key fields:

- `Valid`: Indicates whether the field contains a non-null value.
- `Present`: Indicates whether the field was included in the JSON input.

**Note:** When marshaling to JSON, these types do not fully support the `omitempty` tag. If a field's `Present` field is `false`, the field will still be included in the output JSON with a `null` value.

## How does it work?

Have a for example at the following code:
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/mbe81/jsontype"
)

type Person struct {
	FirstName jsontype.NullString `json:"firstName"`
	LastName  jsontype.NullString `json:"lastName"`
	City      jsontype.NullString `json:"city"`
	Age       jsontype.NullInt    `json:"age"`
}

func main() {
	var p Person
	err := json.Unmarshal([]byte(`{"firstName": "John", "lastName": null, "city": "New York"}`), &p)
	if err != nil {
		panic(err) // This examples must not fail
	}
	fmt.Println(p)
}
```

Running this code will print the following output:

```
{{John true true} { false true} {New York true true} {0 false false}}
```

Which corresponds to the values in the following table:

| Field     | Value        | Valid | Present |
|-----------|--------------|-------|---------|
| FirstName | `"John"`     | true  | true    |
| LastName  | `""`         | false | true    |
| City      | `"New York"` | true  | true    |
| Age       | `0`          | false | false   |

Based on this you can say the following:
- `lastName` is `null` in the JSON because `LastName.Valid` is false and `LastName.Present` is true.
- `age` is absent in the JSON because `Age.Present` is false.

## Supported types

Currently the following types are supported:

- `jsontype.NullString`
- `jsontype.NullInt`
- `jsontype.NullFloat64`
- `jsontype.NullBool`
- `jsontype.NullTime`

## License

This package is released under the MIT license. See the [LICENSE](LICENSE) file for more information. Feel free to use the package as is or copy the types for use in your own projects.

## Contributing

Contributions are welcome! Please feel free to submit a pull request if you find any issues or would like to suggest improvements, for example by adding support for additional types.
