package findutils

import "strings"

// IndexOfString returns the index of string in stringslice
func IndexOfString(strslice []string, val string) int {
	for i, ele := range strslice {
		if ele == val {
			return i
		}
	}
	return -1
}

// IndexOfInt returns the index of int in intslice
func IndexOfInt(intslice []int, val int) int {
	for i, ele := range intslice {
		if ele == val {
			return i
		}
	}
	return -1
}

// IndexOfInt64 returns index of int64 element if present in the slice otherwise -1
func IndexOfInt64(intslice []int64, val int64) int {
	for i, ele := range intslice {
		if ele == val {
			return i
		}
	}
	return -1
}

// IndexOfRune return the index of the rune if present in the runeslice
func IndexOfRune(runeslice []rune, val rune) int {
	for i, ele := range runeslice {
		if ele == val {
			return i
		}
	}
	return -1
}

// ContainsString returns true if string is present in stringslice
func ContainsString(strslice []string, val string) bool {
	return IndexOfString(strslice, val) >= 0
}

// ContainsInt returns true if int is present in intslice
func ContainsInt(intslice []int, val int) bool {
	return IndexOfInt(intslice, val) >= 0
}

// ContainsInt64 returns whether given int64 is present in int64 slice or not
func ContainsInt64(intslice []int64, val int64) bool {
	return IndexOfInt64(intslice, val) >= 0
}

// ContainsRune returns true if the rune is present in the runeslice
func ContainsRune(runeslice []rune, val rune) bool {
	return IndexOfRune(runeslice, val) >= 0
}

// GetStringBefore returns the remaining string before pat.
// If the pat is not found or found at the starting then an empty string is returned
func GetStringBefore(line, pat string) (result string) {
	idx := strings.Index(line, pat)
	if idx >= 0 {
		result = line[:idx]
	}
	return
}

// GetStringAfter returns the remaining string after pat and starting index of remaining string
func GetStringAfter(line, pat string) (result string, idx int) {
	idx = strings.Index(line, pat)
	if idx >= 0 {
		idx += len(pat)
		result = line[idx:]
	}
	return
}

// GetStringBetween returns the string between pat1 and pat2 if not found empty string is returned
// second argument is the index upto which line has been searched
func GetStringBetween(line string, idx int, pat1, pat2 string) (result string, endIndex int) {
	endIndex = -1
	if idx < 0 || idx > len(line) {
		return
	}
	line = line[idx:]
	idx1 := strings.Index(line, pat1)
	if idx1 >= 0 {
		idx2 := idx1 + len(pat1)
		j := strings.Index(line[idx2:], pat2)
		if j >= 0 {
			result = line[idx2 : idx2+j]
			endIndex = idx + idx2 + j + len(pat2)
			return
		}
	}
	return
}

// GetColumnStrictly returns the column after splitting it with a given delim if column is found other returns empty
func GetColumnStrictly(str string, delim string, colidx int) string {
	data := strings.Split(str, delim)
	if colidx < 0 || colidx >= len(data) {
		return ""
	}
	return data[colidx]
}

// GetColumnsStrictly returns selected columns after splitting by separator and joined result by separator. If any of
// the column is not found then empty string is returned
func GetColumnsStrictly(str string, delim string, colindexes []int) string {
	data := strings.Split(str, delim)
	result := []string{}
	for _, colidx := range colindexes {
		if colidx < 0 || colidx >= len(data) {
			return ""
		}
		result = append(result, data[colidx])
	}
	return strings.Join(result, delim)
}
