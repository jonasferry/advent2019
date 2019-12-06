package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel1(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, test := range tests {
		assert.Equal(test.expected, fuel1(test.input))
	}
}

func TestFuel2(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, test := range tests {
		assert.Equal(test.expected, fuel2(test.input))
	}
}
