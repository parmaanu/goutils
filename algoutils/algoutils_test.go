package algoutils_test

import (
	"github.com/parmaanu/goutils/algoutils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	slice := []string{"hello", "how", "are", "you"}
	assert.True(t, algoutils.Any(slice, func(s string) bool {
		return s == "are"
	}))
	assert.False(t, algoutils.Any(slice, func(s string) bool {
		return s == "test"
	}))
}

func TestAll(t *testing.T) {
	slice := []string{"hello", "how", "hare", "hyou"}
	assert.True(t, algoutils.All(slice, func(s string) bool {
		return strings.Contains(s, "h")
	}))
	assert.False(t, algoutils.All(slice, func(s string) bool {
		return strings.Contains(s, "o")
	}))
}

func TestStringContainsAll(t *testing.T) {
	str := "hello how are you?"
	assert.True(t, algoutils.StringContainsAll(str, []string{"hello", "how"}))
	assert.True(t, algoutils.StringContainsAll(str, []string{}))
	assert.False(t, algoutils.StringContainsAll(str, []string{"you", "what"}))
}

func TestFilter(t *testing.T) {
	slice := []string{"hello", "how", "hare", "you"}
	assert.ElementsMatch(t, []string{"hello", "how", "hare"}, algoutils.Filter(slice, func(s string) bool {
		return strings.Contains(s, "h")
	}), "filtered elements are not correct")
	assert.ElementsMatch(t, []string{}, algoutils.Filter(slice, func(s string) bool {
		return strings.Contains(s, "z")
	}), "filtered elements are not correct")
}

func TestMap(t *testing.T) {
	slice := []string{"hello", "how", "hare", "you"}
	assert.ElementsMatch(t, []string{"HELLO", "HOW", "HARE", "YOU"}, algoutils.Map(slice, strings.ToUpper), "filtered elements are not correct")
	assert.ElementsMatch(t, []string{"he", "ho", "ha", "yo"}, algoutils.Map(slice, func(s string) string {
		return s[:2]
	}), "filtered elements are not correct")
}

func TestSlice(t *testing.T) {
	assert.ElementsMatch(t, []string{"he", "ho", "ha", "yo"}, algoutils.Slice("he", "ho", "ha", "yo"))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 4, algoutils.Max(2, 4))
	assert.Equal(t, 4, algoutils.Max(4, 2))
	assert.Equal(t, 4, algoutils.Max(4, 2))
}

func TestMaxIntSlice(t *testing.T) {
	assert.Equal(t, 4, algoutils.MaxIntSlice([]int{2, 1, 4}))
	assert.Equal(t, 4, algoutils.MaxIntSlice([]int{4, 1, 2}))
	assert.Equal(t, 4, algoutils.MaxIntSlice([]int{4, 1, 2}))
}

func TestMaxIntVarible(t *testing.T) {
	assert.Equal(t, 4, algoutils.MaxIntVarible(2, 1, 4))
	assert.Equal(t, 4, algoutils.MaxIntVarible(4, 1, 2))
	assert.Equal(t, 4, algoutils.MaxIntVarible(4, 1, 2))
	assert.Equal(t, 4, algoutils.MaxIntVarible(4))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 2, algoutils.Min(2, 4))
	assert.Equal(t, 2, algoutils.Min(4, 2))
	assert.Equal(t, 2, algoutils.Min(4, 2))
}

func TestMinInt64(t *testing.T) {
	assert.Equal(t, int64(2), algoutils.MinInt64(int64(2), int64(4)))
	assert.Equal(t, int64(2), algoutils.MinInt64(int64(4), int64(2)))
	assert.Equal(t, int64(2), algoutils.MinInt64(int64(4), int64(2)))
}

func TestMinIntSlice(t *testing.T) {
	assert.Equal(t, 1, algoutils.MinIntSlice([]int{2, 1, 4}))
	assert.Equal(t, 1, algoutils.MinIntSlice([]int{4, 1, 2}))
	assert.Equal(t, 1, algoutils.MinIntSlice([]int{4, 1, 2}))
	assert.Equal(t, 1, algoutils.MinIntSlice([]int{1}))
	assert.Equal(t, 0, algoutils.MinIntSlice([]int{}))
}

func TestMinIntVarible(t *testing.T) {
	assert.Equal(t, 1, algoutils.MinIntVarible(2, 1, 4))
	assert.Equal(t, 1, algoutils.MinIntVarible(4, 1, 2))
	assert.Equal(t, 1, algoutils.MinIntVarible(4, 1, 2))
	assert.Equal(t, 1, algoutils.MinIntVarible(1))
}

func TestGetDigitsCount(t *testing.T) {
	assert.Equal(t, 4, algoutils.GetDigitsCount("1234"), "digit count not correct")
	assert.Equal(t, 4, algoutils.GetDigitsCount("a1234"), "digit count not correct")
	assert.Equal(t, 0, algoutils.GetDigitsCount("abc"), "digit count not correct")
}

func TestGetDigitAndCharCount(t *testing.T) {
	assert.Equal(t, 4, algoutils.GetDigitAndCharCount("1234"), "digit count not correct")
	assert.Equal(t, 4, algoutils.GetDigitAndCharCount("a1234"), "digit count not correct")
	assert.Equal(t, 0, algoutils.GetDigitAndCharCount("abc"), "digit count not correct")
	assert.Equal(t, 0, algoutils.GetDigitAndCharCount("abc(", ')'), "digit count not correct")
	assert.Equal(t, 1, algoutils.GetDigitAndCharCount("abc(", '(', ')'), "digit count not correct")
	assert.Equal(t, 2, algoutils.GetDigitAndCharCount("abc()", '(', ')'), "digit count not correct")
	assert.Equal(t, 6, algoutils.GetDigitAndCharCount("<abc>(1234)", '(', ')'), "digit count not correct")
	assert.Equal(t, 6, algoutils.GetDigitAndCharCount("(1234)", '(', ')'), "digit count not correct")
}

func TestSortedKeysString(t *testing.T) {
	m := map[string]interface{}{
		"z": 2,
		"b": 2,
		"y": 2,
		"a": 1,
		"d": 2,
		"c": 2,
	}
	{
		exp := []string{"a", "b", "c", "d", "y", "z"}
		act := algoutils.SortedKeysString(m, false)
		assert.Equal(t, exp, act, "list not sorted")
	}

	{
		exp := []string{"z", "y", "d", "c", "b", "a"}
		act := algoutils.SortedKeysString(m, true)
		assert.Equal(t, exp, act, "list not reverse sorted")
	}
}
