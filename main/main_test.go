package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	t.Run("case 121", func(t *testing.T) {
		result := Palindrome(121)
		assert.True(t, result)
	})

	t.Run("case -121", func(t *testing.T) {
		result := Palindrome(-121)
		assert.False(t, result)
	})

	t.Run("case 10", func(t *testing.T) {
		result := Palindrome(10)
		assert.False(t, result)
	})

	t.Run("case 0", func(t *testing.T) {
		result := Palindrome(0)
		assert.True(t, result)
	})
}

func TestReverse(t *testing.T) {
	t.Run("123", func(t *testing.T) {
		result := Reverse(123)
		assert.Equal(t, 321, result)
	})
	t.Run("121", func(t *testing.T) {
		result := Reverse(121)
		assert.Equal(t, 121, result)
	})
	t.Run("10", func(t *testing.T) {
		result := Reverse(10)
		assert.Equal(t, 1, result)
	})
	t.Run("-121", func(t *testing.T) {
		result := Reverse(-121)
		assert.Equal(t, -121, result)
	})
	t.Run("1", func(t *testing.T) {
		result := Reverse(1)
		assert.Equal(t, 1, result)
	})
	t.Run("0", func(t *testing.T) {
		result := Reverse(0)
		assert.Equal(t, 0, result)
	})
}