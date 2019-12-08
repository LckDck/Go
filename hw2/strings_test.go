package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, "", expand(""))
}
func Test2(t *testing.T) {
	assert.Equal(t, "qwerty", expand("9qwerty"))

}
func Test3(t *testing.T) {
	assert.Equal(t, "qqqqwerty", expand("q4werty"))
}

func Test4(t *testing.T) {
	assert.Equal(t, "qwerty", expand("444qwerty"))
}
func Test5(t *testing.T) {
	assert.Equal(t, "qqqqqqqqqqqwwwerty", expand("q11w3erty"))
}

func Test6(t *testing.T) {
	assert.Equal(t, "qweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerty", expand("qwe100rty"))
}

func Test7(t *testing.T) {
	assert.Equal(t, "qwertyy", expand("qwerty2"))
}

func Test8(t *testing.T) {
	assert.Equal(t, "qwertyyyyyyyyyyyyyyyyyyyyyy", expand("qwerty22"))
}

func Test9(t *testing.T) {
	assert.Equal(t, "qwerty444", expand("qwerty"+"\\"+"43"))
}
func Test10(t *testing.T) {
	assert.Equal(t, "qwerty4", expand("qwerty"+"\\"+"4"))
}
