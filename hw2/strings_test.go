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
	assert.Equal(t, "qwerty444", expand("qwerty\\43"))
}
func Test10(t *testing.T) {
	assert.Equal(t, "qwerty4", expand("qwerty\\4"))
}

func Test11(t *testing.T) {
	assert.Equal(t, "qwerty4yuiop", expand("qwerty\\4yuiop"))
}

func Test12(t *testing.T) {
	assert.Equal(t, "qwertyyyy", expand("\\qwerty4"))
}

func Test13(t *testing.T) {
	assert.Equal(t, "\\qwerty", expand("\\\\qwerty"))
}

func Test14(t *testing.T) {
	assert.Equal(t, "\\\\\\qwerty", expand("\\\\3qwerty"))
}

func Test15(t *testing.T) {
	assert.Equal(t, "qwerty", expand("qwerty\\"))
}

func Test16(t *testing.T) {
	assert.Equal(t, "qwertyyy", expand("qwerty3\\"))
}
