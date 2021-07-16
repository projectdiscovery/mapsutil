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
