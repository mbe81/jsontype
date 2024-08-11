package jsontype

import (
	"testing"
	"time"
)

// Test the UnmarshalJSON method of Bool
func TestBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  Bool
		expectErr bool
	}{
		{
			name:      "Valid bool",
			input:     []byte(`true`),
			expected:  Bool{Bool: true, Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  Bool{Bool: false, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  Bool{Bool: false, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (string)",
			input:     []byte(`"hello"`),
			expected:  Bool{Bool: false, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  Bool{Bool: false, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Bool
			err := b.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if b != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", b, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of Bool
func TestBool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     Bool
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     Bool{Bool: false, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     Bool{Bool: true, Valid: true, Present: true},
			expected:  []byte(`true`),
			expectErr: false,
		},
		{
			name:      "False with both Valid and Present true",
			input:     Bool{Bool: false, Valid: true, Present: true},
			expected:  []byte(`false`),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalJSON()
			if (err != nil) != tt.expectErr {
				t.Errorf("MarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if string(result) != string(tt.expected) {
				t.Errorf("MarshalJSON() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

// Test the UnmarshalJSON method of Float64
func TestFloat64_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  Float64
		expectErr bool
	}{
		{
			name:      "Valid float",
			input:     []byte(`4.56`),
			expected:  Float64{Float64: 4.56, Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  Float64{Float64: 0, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  Float64{Float64: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (string)",
			input:     []byte(`"hello"`),
			expected:  Float64{Float64: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  Float64{Float64: 0, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f Float64
			err := f.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if f != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", f, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of Float64
func TestFloat64_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     Float64
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     Float64{Float64: 0, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     Float64{Float64: 4.56, Valid: true, Present: true},
			expected:  []byte(`4.56`),
			expectErr: false,
		},
		{
			name:      "Zero with both Valid and Present true",
			input:     Float64{Float64: 0, Valid: true, Present: true},
			expected:  []byte(`0`),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalJSON()
			if (err != nil) != tt.expectErr {
				t.Errorf("MarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if string(result) != string(tt.expected) {
				t.Errorf("MarshalJSON() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

// Test the UnmarshalJSON method of Int
func TestInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  Int
		expectErr bool
	}{
		{
			name:      "Valid int",
			input:     []byte(`123`),
			expected:  Int{Int: 123, Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  Int{Int: 0, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  Int{Int: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (string)",
			input:     []byte(`"hello"`),
			expected:  Int{Int: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (float)",
			input:     []byte(`"4.56"`),
			expected:  Int{Int: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  Int{Int: 0, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i Int
			err := i.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if i != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", i, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of Int
func TestInt_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     Int
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     Int{Int: 0, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     Int{Int: 123, Valid: true, Present: true},
			expected:  []byte(`123`),
			expectErr: false,
		},
		{
			name:      "Zero with both Valid and Present true",
			input:     Int{Int: 0, Valid: true, Present: true},
			expected:  []byte(`0`),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalJSON()
			if (err != nil) != tt.expectErr {
				t.Errorf("MarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if string(result) != string(tt.expected) {
				t.Errorf("MarshalJSON() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

// Test the UnmarshalJSON method of String
func TestString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  String
		expectErr bool
	}{
		{
			name:      "Valid string",
			input:     []byte(`"hello"`),
			expected:  String{String: "hello", Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  String{String: "", Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  String{String: "", Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (number)",
			input:     []byte(`123`),
			expected:  String{String: "", Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  String{String: "", Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s String
			err := s.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if s != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", s, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of String
func TestString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     String
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     String{String: "", Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     String{String: "hello", Valid: true, Present: true},
			expected:  []byte(`"hello"`),
			expectErr: false,
		},
		{
			name:      "Empty string with both Valid and Present true",
			input:     String{String: "", Valid: true, Present: true},
			expected:  []byte(`""`),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalJSON()
			if (err != nil) != tt.expectErr {
				t.Errorf("MarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if string(result) != string(tt.expected) {
				t.Errorf("MarshalJSON() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

// Test the UnmarshalJSON method of String
func TestTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  Time
		expectErr bool
	}{
		{
			name:  "Valid string",
			input: []byte(`"2024-01-01T00:00:00Z"`),
			expected: Time{
				Time:    time.Date(2024, 01, 01, 00, 00, 00, 00, time.UTC),
				Valid:   true,
				Present: true,
			},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  Time{Time: time.Time{}, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  Time{Time: time.Time{}, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (number)",
			input:     []byte(`123`),
			expected:  Time{Time: time.Time{}, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  Time{Time: time.Time{}, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Time
			err := got.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if got != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of String
func TestTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     Time
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     Time{Time: time.Time{}, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name: "Both Valid and Present are true",
			input: Time{
				Time:    time.Date(2024, 01, 01, 00, 00, 00, 00, time.UTC),
				Valid:   true,
				Present: true,
			},
			expected:  []byte(`"2024-01-01T00:00:00Z"`),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalJSON()
			if (err != nil) != tt.expectErr {
				t.Errorf("MarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if string(result) != string(tt.expected) {
				t.Errorf("MarshalJSON() = %s, expected %s", result, tt.expected)
			}
		})
	}
}
