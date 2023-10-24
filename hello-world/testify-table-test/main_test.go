package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	testHelloCases := []struct {
		expected string
	}{
		{"Hello World!"},
		// Here in another examples you just need to put more
		// strings or ints to test more possibilities. You can pass more variables from your struct too
		// if you have more arguments like a Sum of two numbers.
	}

	for _, testCases := range testHelloCases {
		t.Run("Test Hello World message", func(t *testing.T) {
			result := HelloWorld()
			assert.Equal(t, testCases.expected, result)
		})
	}

}
