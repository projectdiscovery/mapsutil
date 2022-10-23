package mapsutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeMaps(t *testing.T) {
	m1Str := map[string]interface{}{"a": 1, "b": 2}
	m2Str := map[string]interface{}{"b": 2, "c": 3}
	rStr := map[string]interface{}{"a": 1, "b": 2, "c": 3}
	rrStr := MergeMaps(m1Str, m2Str)
	require.EqualValues(t, rStr, rrStr)

	m1Int := map[int]interface{}{1: 1, 2: 2}
	m2Int := map[int]interface{}{1: 1, 2: 2, 3: 3, 4: 4}
	m3Int := map[int]interface{}{1: 1, 5: 5}
	rInt := map[int]interface{}{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	rrInt := MergeMaps(m1Int, m2Int, m3Int)
	require.EqualValues(t, rInt, rrInt)
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
