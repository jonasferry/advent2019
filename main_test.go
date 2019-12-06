package main

import (
	"testing"

	"github.com/jonasferry/advent2019/day01"
	"github.com/jonasferry/advent2019/day02"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3318604, day01.Run1())
	assert.Equal(4975039, day01.Run2())
}

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(4330636, day02.Run1())
}
