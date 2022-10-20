package mapsutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeMaps(t *testing.T) {
	m1 := map[string]interface{}{"a": 1, "b": 2}
	m2 := map[string]interface{}{"b": 2, "c": 3}
	r := map[string]interface{}{"a": 1, "b": 2, "c": 3}
	rr := MergeMaps(m1, m2)
	require.Equal(t, r, rr)
}

func TestMergeMapsWithStrings(t *testing.T) {
	m1 := map[string]string{"a": "a", "b": "b"}
	m2 := map[string]string{"b": "b", "c": "c"}
	r := map[string]string{"a": "a", "b": "b", "c": "c"}
	rr := MergeMapsWithStrings(m1, m2)
	require.Equal(t, r, rr)
}

func TestHTTPToMap(t *testing.T) {
	// not implemented
}

func TestDNSToMap(t *testing.T) {
	// not implemented
}

func TestHTTPRequesToMap(t *testing.T) {
	// not implemented
}

func TestHTTPResponseToMap(t *testing.T) {
	// not implemented
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		inputA map[string]interface{}
		inputB map[string]interface{}
		want   map[string]interface{}
	}{
		{
			name:   "single value",
			inputA: map[string]interface{}{"a": "301"},
			inputB: map[string]interface{}{},
			want:   map[string]interface{}{"a": "301"},
		},
		{
			name:   "empty value",
			inputA: map[string]interface{}{"a": 1},
			inputB: map[string]interface{}{"": ""},
			want:   map[string]interface{}{"a": 1},
		},
		{
			name:   "empty map",
			inputA: map[string]interface{}{},
			inputB: map[string]interface{}{},
			want:   map[string]interface{}{},
		},
		{
			name:   "different values",
			inputA: map[string]interface{}{"a": true, "b": false, "c": true, "d": false},
			inputB: map[string]interface{}{"c": "C", "d": "D"},
			want:   map[string]interface{}{"a": true, "b": false},
		},
		{
			name:   "differenceSTR",
			inputA: map[string]interface{}{"a": "A", "b": "B", "c": "C", "d": "D"},
			inputB: map[string]interface{}{"c": "C", "d": "D"},
			want:   map[string]interface{}{"a": "A", "b": "B"},
		},
		{
			name:   "differenceINT",
			inputA: map[string]interface{}{"a": 1, "b": 2, "c": 3, "d": 4},
			inputB: map[string]interface{}{"c": 3, "d": 4},
			want:   map[string]interface{}{"a": 1, "b": 2},
		},
		{
			name:   "differenceBOOL",
			inputA: map[string]interface{}{"a": true, "b": false, "c": true, "d": false},
			inputB: map[string]interface{}{"c": true, "d": false},
			want:   map[string]interface{}{"a": true, "b": false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.inputA, tt.inputB)
			require.Equal(t, tt.want, got)
		})
	}
}
