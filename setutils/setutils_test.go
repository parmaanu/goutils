package setutils_test

import (
	"github.com/parmaanu/goutils/setutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendIfMissingForStringSet(t *testing.T) {
	set := setutils.StringSet{"hello", "how", "are"}
	expectedSet := []string{"hello", "how", "are"}
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing("vivek")
	expectedSet = append(expectedSet, "vivek")
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing("how")
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing("I", "am", "good")
	expectedSet = append(expectedSet, "I", "am", "good")
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	// check ordering of elements
	assert.Equal(t, "hello", set[0], "set ordering is changed")
	assert.Equal(t, "are", set[2], "set ordering is changed")
	assert.Equal(t, "am", set[5], "set ordering is changed")
}

func TestAppendIfMissingForIntSet(t *testing.T) {
	set := setutils.IntSet{11, 22, 33}
	expectedSet := []int{11, 22, 33}
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing(44)
	expectedSet = append(expectedSet, 44)
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing(22)
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing(44, 55, 66)
	expectedSet = append(expectedSet, 55, 66)
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	// check ordering of elements
	assert.Equal(t, 11, set[0], "set ordering is changed")
	assert.Equal(t, 33, set[2], "set ordering is changed")
	assert.Equal(t, 66, set[5], "set ordering is changed")
}

func TestAppendIfMissingForInt64Set(t *testing.T) {
	set := setutils.Int64Set{11, 22, 33}
	expectedSet := []int64{11, 22, 33}
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing(44)
	expectedSet = append(expectedSet, 44)
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing(22)
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	set = set.AppendIfMissing(44, 55, 66)
	expectedSet = append(expectedSet, 55, 66)
	assert.ElementsMatch(t, expectedSet, set, "set not equal")

	// check ordering of elements
	assert.Equal(t, int64(11), set[0], "set ordering is changed")
	assert.Equal(t, int64(33), set[2], "set ordering is changed")
	assert.Equal(t, int64(66), set[5], "set ordering is changed")
}

func TestUniqueStringSlice(t *testing.T) {
	strSlice := []string{"hello", "how", "are", "you", "are", "you", "fine"}
	uniqueSlice := setutils.UniqueStringSlice(strSlice)
	expectedSlice := []string{"hello", "how", "are", "you", "fine"}
	assert.Equal(t, expectedSlice, uniqueSlice, "UniqueStringSlice does not return unique slices of strings")
}
