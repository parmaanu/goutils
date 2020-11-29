package algoutils

import (
	"github.com/parmaanu/goutils/findutils"
	"hash/fnv"
	"sort"
	"strings"
)

// Any returns true if one of the strings in the slice satisfies the predicate f.
// Example: fmt.Println(Any(strs, func(v string) bool {
// return strings.HasPrefix(v, "p")
// }))
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all of the strings in the slice satisfy the predicate f.
// Example: fmt.Println(All(strs, func(v string) bool {
// return strings.HasPrefix(v, "p")
// }))
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// StringContainsAll returns true when all the patterns are present in the string 'line'
func StringContainsAll(line string, patterns []string) bool {
	return All(patterns, func(str string) bool {
		return strings.Contains(line, str)
	})
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
// Example: fmt.Println(Filter(strs, func(v string) bool {
// return strings.Contains(v, "e")
// }))
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function f to each string in the original slice.
// Example: fmt.Println(Map(strs, strings.ToUpper))
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Slice returns a slice out of variadic arguments
func Slice(args ...interface{}) []interface{} {
	return args
}

// Max returns the maximum of two ints
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MaxIntSlice returns the maximum from an int slice
func MaxIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = MaxIntVarible(v[0], v[1:]...)
	}
	return
}

// MaxIntVarible returns the maximum from variadic arguments
func MaxIntVarible(v1 int, vn ...int) (m int) {
	m = v1
	for i := 0; i < len(vn); i++ {
		if vn[i] > m {
			m = vn[i]
		}
	}
	return
}

// Min returns the minimum of two variables
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MinInt64 return the minimum of two Int64
func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// MinIntSlice returns the minimum from int slice. Zero is returned if slice is empty
func MinIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = MinIntVarible(v[0], v[1:]...)
	}
	return
}

// MinIntVarible returns the minimum from variadic arguments
func MinIntVarible(v1 int, vn ...int) (m int) {
	m = v1
	for i := 0; i < len(vn); i++ {
		if vn[i] < m {
			m = vn[i]
		}
	}
	return
}

// GetDigitsCount counts the number of digits in the string
func GetDigitsCount(str string) int {
	cnt := 0
	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			cnt++
		}
	}
	return cnt
}

// GetDigitAndCharCount returns the count of char and digits
func GetDigitAndCharCount(str string, val ...rune) int {
	cnt := 0
	for _, ch := range str {
		if (ch >= '0' && ch <= '9') || findutils.ContainsRune(val, ch) {
			cnt++
		}
	}
	return cnt
}

// SortedKeysString returns sorted string keys
func SortedKeysString(m map[string]interface{}, reverse bool) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}
	return keys
}

// GetHash32 return the hash for a stringslice
func GetHash32(strslice []string) uint32 {
	h := fnv.New32a()
	for _, s := range strslice {
		h.Write([]byte(s))
	}
	return h.Sum32()
}
