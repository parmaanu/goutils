package setutils

// IntSet is a set of integers
type IntSet []int

// AppendIfMissing appends to int slice only if it is missing
func (slice IntSet) AppendIfMissing(val ...int) IntSet {
	m := make(map[int]bool, len(slice))
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

// Int64Set is a set of int64
type Int64Set []int64

// AppendIfMissing appends to the slice only if it is missing
func (slice Int64Set) AppendIfMissing(val ...int64) Int64Set {
	m := make(map[int64]bool, len(slice))
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
