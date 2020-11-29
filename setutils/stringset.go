package setutils

// StringSet is a set of string
type StringSet []string

// AppendIfMissing appends into string set only if the element is missing
// may not be best efficient (there's one more function call O(n^2)
func (slice StringSet) AppendIfMissing(val ...string) StringSet {
	m := make(map[string]bool, len(slice))
	for _, ele := range slice {
		m[ele] = true
	}
	for _, ele := range val {
		_, exists := m[ele]
		if !exists {
			slice = append(slice, ele)
		}
	}
	return slice
}

// UniqueStringSlice returns a string slice after removing duplicate elements
func UniqueStringSlice(input []string) []string {
	u := []string{}
	m := make(map[string]bool)
	for _, val := range input {
		if _, exists := m[val]; !exists {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
