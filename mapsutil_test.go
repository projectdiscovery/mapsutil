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

func TestGetKeys(t *testing.T) {
	t.Run("GetKeys(empty)", func(t *testing.T) {
		got := GetKeys(map[string]interface{}{})
		require.EqualValues(t, []string{}, got)
	})

	t.Run("GetKeys(string)", func(t *testing.T) {
		got := GetKeys(map[string]interface{}{"a": "a", "b": "b"})
		require.EqualValues(t, []string{"a", "b"}, got)
	})

	t.Run("GetKeys(int)", func(t *testing.T) {
		got := GetKeys(map[int]interface{}{1: "a", 2: "b"})
		require.EqualValues(t, []int{1, 2}, got)
	})

	t.Run("GetKeys(bool)", func(t *testing.T) {
		got := GetKeys(map[bool]interface{}{true: "a", false: "b"})
		require.EqualValues(t, []bool{true, false}, got)
	})
}

func TestDifference(t *testing.T) {
	t.Run("Difference(empty)", func(t *testing.T) {
		got := Difference(map[string]interface{}{}, []string{}...)
		require.EqualValues(t, map[string]interface{}{}, got)
	})

	t.Run("Difference(string)", func(t *testing.T) {
		got := Difference(map[string]interface{}{"a": 1, "b": 2, "c": 3}, []string{"a"}...)
		require.EqualValues(t, map[string]interface{}{"b": 2, "c": 3}, got)
	})

	t.Run("Difference(int)", func(t *testing.T) {
		got := Difference(map[int]interface{}{1: "a", 2: "b", 3: "c"}, []int{1}...)
		require.EqualValues(t, map[int]interface{}{2: "b", 3: "c"}, got)
	})

	t.Run("Difference(bool)", func(t *testing.T) {
		got := Difference(map[bool]interface{}{true: 1, false: 2}, []bool{true}...)
		require.EqualValues(t, map[bool]interface{}{false: 2}, got)
	})
}
