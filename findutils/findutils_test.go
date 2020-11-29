package findutils_test

import (
	"github.com/parmaanu/goutils/findutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfString(t *testing.T) {
	slice := []string{"hello", "how", "are", "you?"}
	assert.Equal(t, 2, findutils.IndexOfString(slice, "are"), "incorrect index while finding string in a string slice")
	assert.Equal(t, -1, findutils.IndexOfString(slice, "aree"), "incorrect index while finding string in a string slice")
}

func TestIndexOfInt(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	assert.Equal(t, 2, findutils.IndexOfInt(slice, 3), "incorrect index while finding int in a int slice")
	assert.Equal(t, -1, findutils.IndexOfInt(slice, 5), "incorrect index while finding int in a int slice")
}

func TestIndexOfInt64(t *testing.T) {
	slice := []int64{1, 2, 3, 4}
	assert.Equal(t, 2, findutils.IndexOfInt64(slice, 3), "incorrect index while finding int64 in a int64 slice")
	assert.Equal(t, -1, findutils.IndexOfInt64(slice, 5), "incorrect index while finding int64 in a int64 slice")
}

func TestIndexOfRune(t *testing.T) {
	slice := []rune{'1', '2', '3', '4'}
	assert.Equal(t, 2, findutils.IndexOfRune(slice, '3'), "incorrect index while finding rune in a rune slice")
	assert.Equal(t, -1, findutils.IndexOfRune(slice, '5'), "incorrect index while finding rune in a rune slice")
}

func TestContainsString(t *testing.T) {
	slice := []string{"hello", "how", "are", "you?"}
	assert.True(t, findutils.ContainsString(slice, "are"), "string should exist in slice")
	assert.False(t, findutils.ContainsString(slice, "aree"), "string should not exist in slice")
}

func TestContainsInt(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	assert.True(t, findutils.ContainsInt(slice, 3), "int should exist in slice")
	assert.False(t, findutils.ContainsInt(slice, 5), "int should not exist in slice")
}

func TestContainsInt64(t *testing.T) {
	slice := []int64{1, 2, 3, 4}
	assert.True(t, findutils.ContainsInt64(slice, 3), "int64 should exist in slice")
	assert.False(t, findutils.ContainsInt64(slice, 5), "int64 should not exist in slice")
}

func TestContainsRune(t *testing.T) {
	slice := []rune{'1', '2', '3', '4'}
	assert.True(t, findutils.ContainsRune(slice, '3'), "rune should exist in slice")
	assert.False(t, findutils.ContainsRune(slice, '5'), "run should not exist in slice")
}

func TestGetStringBefore(t *testing.T) {
	line := "this is a very big logline"
	assert.Equal(t, "this is", findutils.GetStringBefore(line, " a "), "GetStringBefore does not return string before pattern")
	assert.Equal(t, "", findutils.GetStringBefore(line, " aasdf "), "GetStringBefore does not return string before pattern")
}

func TestGetStringAfter(t *testing.T) {
	line := "this is a very big logline"
	{
		str, idx := findutils.GetStringAfter(line, " a ")
		assert.Equal(t, "very big logline", str, "GetStringAfter does not return a correct string after the pattern")
		assert.Equal(t, 10, idx, "GetStringAfter does not correct index after of the string after the pattern")
	}
	{
		str, idx := findutils.GetStringAfter(line, " aasdf ")
		assert.Equal(t, "", str, "GetStringAfter does not return a correct string after the pattern")
		assert.Equal(t, -1, idx, "GetStringAfter does not correct index after of the string after the pattern")
	}
}

func TestGetStringBetween(t *testing.T) {
	line := "client_order_id: 39338 common, test: testing, quantity: 10000, price: 12312.23, bid[123.0]"
	// client_order_id:[16] 39338[22] common,[30] test:[36] testing,[45] quantity:[55] 10000,[62] price:[69]
	// 12312.23,[79] bid[123.0][91]"
	{
		orderidstr, idx := findutils.GetStringBetween(line, 0, "client_order_id: ", " ")
		assert.Equal(t, "39338", orderidstr, "GetStringBetween does not extract string properly")
		assert.Equal(t, 23, idx, "incorrect index returned by GetStringBetween")
	}
	{
		bidstr, idx := findutils.GetStringBetween(line, 0, "bid[", "]")
		assert.Equal(t, "123.0", bidstr, "GetStringBetween does not extract string properly")
		assert.Equal(t, 90, idx, "incorrect index returned by GetStringBetween")
	}
	{
		orderID, idx := findutils.GetStringBetween(line, 0, "order_id: ", ",")
		assert.Equal(t, "39338 common", orderID, "GetStringBetween does not extract string properly")
		assert.Equal(t, 30, idx, "incorrect index returned by GetStringBetween")

		test, idx := findutils.GetStringBetween(line, idx, "test: ", ",")
		assert.Equal(t, "testing", test, "GetStringBetween does not extract string properly")
		assert.Equal(t, 45, idx, "incorrect index returned by GetStringBetween")

		quantity, idx := findutils.GetStringBetween(line, idx, "quantity: ", ",")
		assert.Equal(t, "10000", quantity, "GetStringBetween does not extract string properly")
		assert.Equal(t, 62, idx, "incorrect index returned by GetStringBetween")

		price, idx := findutils.GetStringBetween(line, idx, "price: ", ",")
		assert.Equal(t, "12312.23", price, "GetStringBetween does not extract string properly")
		assert.Equal(t, 79, idx, "incorrect index returned by GetStringBetween")
	}
}

func TestGetColumnStrictly(t *testing.T) {
	str := "order,price,side,quantity,fillprice,fillquantity"
	assert.Equal(t, "order", findutils.GetColumnStrictly(str, ",", 0), "GetColumnStrictly returned incorrect column")
	assert.Equal(t, "price", findutils.GetColumnStrictly(str, ",", 1), "GetColumnStrictly returned incorrect column")
	assert.Equal(t, "", findutils.GetColumnStrictly(str, ",", 60), "GetColumnStrictly returned incorrect column")
	assert.Equal(t, "", findutils.GetColumnStrictly(str, ",", -1), "GetColumnStrictly returned incorrect column")
}

func TestGetColumnsStrictly(t *testing.T) {
	str := "order,price,side,quantity,fillprice,fillquantity"
	assert.Equal(t, "order,side,quantity", findutils.GetColumnsStrictly(str, ",", []int{0, 2, 3}), "GetColumnsStrictly returned incorrect column")
	assert.Equal(t, "", findutils.GetColumnsStrictly(str, ",", []int{1, 4, 60}), "GetColumnsStrictly returned incorrect column")
	assert.Equal(t, "", findutils.GetColumnsStrictly(str, ",", []int{}), "GetColumnsStrictly returned incorrect column")
}
