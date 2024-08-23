package jsontype

import (
	"testing"
	"time"
)

// Test the UnmarshalJSON method of NullBool
func TestNullBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  NullBool
		expectErr bool
	}{
		{
			name:      "Valid bool",
			input:     []byte(`true`),
			expected:  NullBool{Bool: true, Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  NullBool{Bool: false, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  NullBool{Bool: false, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (string)",
			input:     []byte(`"hello"`),
			expected:  NullBool{Bool: false, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  NullBool{Bool: false, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nb NullBool
			err := nb.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if nb != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", nb, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of NullBool
func TestNullBool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     NullBool
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     NullBool{Bool: false, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     NullBool{Bool: true, Valid: true, Present: true},
			expected:  []byte(`true`),
			expectErr: false,
		},
		{
			name:      "False with both Valid and Present true",
			input:     NullBool{Bool: false, Valid: true, Present: true},
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

// Test the UnmarshalJSON method of NullFloat64
func TestNullFloat64_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  NullFloat64
		expectErr bool
	}{
		{
			name:      "Valid float",
			input:     []byte(`4.56`),
			expected:  NullFloat64{Float64: 4.56, Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  NullFloat64{Float64: 0, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  NullFloat64{Float64: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (string)",
			input:     []byte(`"hello"`),
			expected:  NullFloat64{Float64: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  NullFloat64{Float64: 0, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nf NullFloat64
			err := nf.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if nf != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", nf, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of NullFloat64
func TestNullFloat64_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     NullFloat64
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     NullFloat64{Float64: 0, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     NullFloat64{Float64: 4.56, Valid: true, Present: true},
			expected:  []byte(`4.56`),
			expectErr: false,
		},
		{
			name:      "Zero with both Valid and Present true",
			input:     NullFloat64{Float64: 0, Valid: true, Present: true},
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

// Test the UnmarshalJSON method of NullInt
func TestNullInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  NullInt
		expectErr bool
	}{
		{
			name:      "Valid int",
			input:     []byte(`123`),
			expected:  NullInt{Int: 123, Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  NullInt{Int: 0, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  NullInt{Int: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (string)",
			input:     []byte(`"hello"`),
			expected:  NullInt{Int: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (float)",
			input:     []byte(`"4.56"`),
			expected:  NullInt{Int: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  NullInt{Int: 0, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ni NullInt
			err := ni.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if ni != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", ni, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of NullInt
func TestNullInt_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     NullInt
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     NullInt{Int: 0, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     NullInt{Int: 123, Valid: true, Present: true},
			expected:  []byte(`123`),
			expectErr: false,
		},
		{
			name:      "Zero with both Valid and Present true",
			input:     NullInt{Int: 0, Valid: true, Present: true},
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

// Test the UnmarshalJSON method of NullString
func TestNullString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  NullString
		expectErr bool
	}{
		{
			name:      "Valid string",
			input:     []byte(`"hello"`),
			expected:  NullString{String: "hello", Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  NullString{String: "", Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  NullString{String: "", Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (number)",
			input:     []byte(`123`),
			expected:  NullString{String: "", Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  NullString{String: "", Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ns NullString
			err := ns.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if ns != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", ns, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of NullString
func TestNullString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     NullString
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     NullString{String: "", Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     NullString{String: "hello", Valid: true, Present: true},
			expected:  []byte(`"hello"`),
			expectErr: false,
		},
		{
			name:      "Empty string with both Valid and Present true",
			input:     NullString{String: "", Valid: true, Present: true},
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

// Test the UnmarshalJSON method of NullTime
func TestNullTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  NullTime
		expectErr bool
	}{
		{
			name:  "Valid time",
			input: []byte(`"2024-01-01T00:00:00Z"`),
			expected: NullTime{
				Time:    time.Date(2024, 01, 01, 00, 00, 00, 00, time.UTC),
				Valid:   true,
				Present: true,
			},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  NullTime{Time: time.Time{}, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  NullTime{Time: time.Time{}, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (number)",
			input:     []byte(`123`),
			expected:  NullTime{Time: time.Time{}, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type (object)",
			input:     []byte(`{}`),
			expected:  NullTime{Time: time.Time{}, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nt NullTime
			err := nt.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if nt != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", nt, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of NullTime
func TestNullTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     NullTime
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     NullTime{Time: time.Time{}, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name: "Both Valid and Present are true",
			input: NullTime{
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

// Test the UnmarshalJSON method of Null[string]
func TestGenericNullString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  Null[string]
		expectErr bool
	}{
		{
			name:      "Valid type",
			input:     []byte(`"hello"`),
			expected:  Null[string]{Value: "hello", Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  Null[string]{Value: "", Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  Null[string]{Value: "", Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type: number",
			input:     []byte(`123`),
			expected:  Null[string]{Value: "", Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type: object",
			input:     []byte(`{}`),
			expected:  Null[string]{Value: "", Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ns Null[string]
			err := ns.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if ns != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", ns, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of Null[string]
func TestGenericNullString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     Null[string]
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     Null[string]{Value: "", Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     Null[string]{Value: "hello", Valid: true, Present: true},
			expected:  []byte(`"hello"`),
			expectErr: false,
		},
		{
			name:      "Empty string with both Valid and Present true",
			input:     Null[string]{Value: "", Valid: true, Present: true},
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

// Test the UnmarshalJSON method of Null[int]
func TestGenericNullInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  Null[int]
		expectErr bool
	}{
		{
			name:      "Valid type: string to string",
			input:     []byte(`123`),
			expected:  Null[int]{Value: 123, Valid: true, Present: true},
			expectErr: false,
		},
		{
			name:      "Null value",
			input:     []byte(`null`),
			expected:  Null[int]{Value: 0, Valid: false, Present: true},
			expectErr: false,
		},
		{
			name:      "Missing field",
			input:     nil, // Simulates missing field (should not unmarshal)
			expected:  Null[int]{Value: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type: string",
			input:     []byte(`"hello"`),
			expected:  Null[int]{Value: 0, Valid: false, Present: false},
			expectErr: true,
		},
		{
			name:      "Invalid type: object",
			input:     []byte(`{}`),
			expected:  Null[int]{Value: 0, Valid: false, Present: false},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ns Null[int]
			err := ns.UnmarshalJSON(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("UnmarshalJSON() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if ns != tt.expected {
				t.Errorf("UnmarshalJSON() = %v, expected %v", ns, tt.expected)
			}
		})
	}
}

// Test the MarshalJSON method of Null[string]
func TestGenericNullInt_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     Null[int]
		expected  []byte
		expectErr bool
	}{
		{
			name:      "Valid is false and Present is true",
			input:     Null[int]{Value: 0, Valid: false, Present: true},
			expected:  []byte("null"),
			expectErr: false,
		},
		{
			name:      "Both Valid and Present are true",
			input:     Null[int]{Value: 123, Valid: true, Present: true},
			expected:  []byte(`123`),
			expectErr: false,
		},
		{
			name:      "Zero value with both Valid and Present true",
			input:     Null[int]{Value: 0, Valid: true, Present: true},
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
