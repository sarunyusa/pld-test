package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindome(t *testing.T) {
	t.Run("case 121", func(t *testing.T) {
		result := Palindome(121)
		assert.True(t, result)
	})

	t.Run("case -121", func(t *testing.T) {
		result := Palindome(-121)
		assert.False(t, result)
	})

	t.Run("case 10", func(t *testing.T) {
		result := Palindome(10)
		assert.False(t, result)
	})
}

func TestReverse(t *testing.T) {
	t.Run("123", func(t *testing.T) {
		result := Reverse(123)
		assert.Equal(t, result, 321)
	})
	t.Run("121", func(t *testing.T) {
		result := Reverse(121)
		assert.Equal(t, result, 121)
	})
	t.Run("10", func(t *testing.T) {
		result := Reverse(10)
		assert.Equal(t, result, 01)
	})
	t.Run("-121", func(t *testing.T) {
		result := Reverse(-121)
		assert.Equal(t, result, -121)
	})
	t.Run("1", func(t *testing.T) {
		result := Reverse(1)
		assert.Equal(t, result, 1)
	})
}