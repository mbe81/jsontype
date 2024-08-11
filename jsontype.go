// Package jsontype provides types for handling JSON unmarshaling in scenarios where fields may be null or absent,
// such as when processing data from PATCH requests.
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
	"encoding/json"
	"time"
)

// Bool represents a bool that may be null or may be absent.
// Bool implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type Bool struct {
	Bool    bool
	Valid   bool // Valid is true if Bool is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		b.Bool, b.Valid, b.Present = false, false, true
		return nil
	}
	b.Bool, b.Valid, b.Present = *value, true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Present || !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Bool)
}

// Float64 represents a float64 that may be null or may be absent.
// Float64 implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type Float64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (f *Float64) UnmarshalJSON(data []byte) error {
	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		f.Float64, f.Valid, f.Present = 0, false, true
		return nil
	}
	f.Float64, f.Valid, f.Present = *value, true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (f Float64) MarshalJSON() ([]byte, error) {
	if !f.Present || !f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.Float64)
}

// Int represents an int that may be null or may be absent.
// Int implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type Int struct {
	Int     int
	Valid   bool // Valid is true if Int is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Int) UnmarshalJSON(data []byte) error {
	var value *int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		i.Int, i.Valid, i.Present = 0, false, true
		return nil
	}
	i.Int, i.Valid, i.Present = *value, true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (i Int) MarshalJSON() ([]byte, error) {
	if !i.Present || !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int)
}

// String represents a string that may be null or may be absent.
// String implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type String struct {
	String  string
	Valid   bool // Valid is true if String is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *String) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		s.String, s.Valid, s.Present = "", false, true
		return nil
	}
	s.String, s.Valid, s.Present = *value, true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Present || !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

// Time represents a time.Time that may be null or may be absent.
// Time implements the json.Unmarshaler and can be used as a json.Unmarshal destination.
type Time struct {
	Time    time.Time
	Valid   bool // Valid is true if String is not NULL
	Present bool // Present is true if the field is present during Unmarshal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(data []byte) error {
	var value *time.Time
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Time, t.Valid, t.Present = time.Time{}, false, true
		return nil
	}
	t.Time, t.Valid, t.Present = *value, true, true
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	if !t.Present || !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time)
}
