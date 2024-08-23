// Package jsontype provides types for handling JSON unmarshaling in scenarios where fields may be null or absent,
// such as when processing data from PATCH requests. The types are inspired by the sql.Null* types in the database/sql package.
//
// The types in this package implement the json.Unmarshaler and json.Marshaler interfaces, allowing them to seamlessly
// integrate with Go's standard JSON handling. Each type includes a Valid field, which indicates whether the value
// is not null, and a Present field, which indicates whether the field was included in the JSON input.
//
// Note: When marshaling, these types do not fully support the 'omitempty' tag. If a field's Present field is false,
// the field will still be included in the output JSON with a null value.
//
// The jsontype package is designed to work alongside the standard library's encoding/json package and should be
// compatible with third-party JSON encoding packages, although this has not been extensively tested.
package jsontype

import (
	"bytes"
	"encoding/json"
	"time"
)

var nullLiteral = []byte("null")

// NullBool represents a bool that may be null or may be absent.
// NullBool implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type NullBool struct {
	Bool    bool
	Valid   bool // Valid is true if Bool is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (nb *NullBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullLiteral) {
		nb.Bool, nb.Valid, nb.Present = false, false, true
		return nil
	}
	if err := json.Unmarshal(data, &nb.Bool); err != nil {
		return err
	}
	nb.Valid, nb.Present = true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Present || !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// NullFloat64 represents a float64 that may be null or may be absent.
// NullFloat64 implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type NullFloat64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullLiteral) {
		nf.Float64, nf.Valid, nf.Present = 0, false, true
		return nil
	}
	if err := json.Unmarshal(data, &nf.Float64); err != nil {
		return err
	}
	nf.Valid, nf.Present = true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Present || !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// NullInt represents an int that may be null or may be absent.
// NullInt implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type NullInt struct {
	Int     int
	Valid   bool // Valid is true if Int is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ni *NullInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullLiteral) {
		ni.Int, ni.Valid, ni.Present = 0, false, true
		return nil
	}
	if err := json.Unmarshal(data, &ni.Int); err != nil {
		return err
	}
	ni.Valid, ni.Present = true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Present || !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int)
}

// NullString represents a string that may be null or may be absent.
// NullString implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type NullString struct {
	String  string
	Valid   bool // Valid is true if String is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ns *NullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullLiteral) {
		ns.String, ns.Valid, ns.Present = "", false, true
		return nil
	}
	if err := json.Unmarshal(data, &ns.String); err != nil {
		return err
	}
	ns.Valid, ns.Present = true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Present || !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// NullTime represents a time.Time that may be null or may be absent.
// NullTime implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type NullTime struct {
	Time    time.Time
	Valid   bool // Valid is true if NullString is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullLiteral) {
		nt.Time, nt.Valid, nt.Present = time.Time{}, false, true
		return nil
	}
	if err := json.Unmarshal(data, &nt.Time); err != nil {
		return err
	}
	nt.Valid, nt.Present = true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Present || !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}

// Null represents a generic value that may be null or may be absent.
// Null implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type Null[T any] struct {
	Value   T
	Valid   bool
	Present bool
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (nt *Null[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nt.Present = true
		return nil
	}
	if err := json.Unmarshal(data, &nt.Value); err != nil {
		return err
	}
	nt.Valid, nt.Present = true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (nt Null[T]) MarshalJSON() ([]byte, error) {
	if !nt.Present || !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Value)
}
