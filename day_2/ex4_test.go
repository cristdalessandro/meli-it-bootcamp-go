package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	avgExp, maxExp, minExp := 19.333333333333332, 64.0, 1.0

	avgOp, _ := operation("average")
	minOp, _ := operation("minimum")
	maxOp, _ := operation("maximum")
	_, ok := operation("_")

	avgRes := avgOp(2, 43, 64, 3, 1, 3)
	minRes := minOp(2, 43, 64, 3, 1, 3)
	maxRes := maxOp(2, 43, 64, 3, 1, 3)

	assert.Equal(t, avgRes, avgExp)
	assert.Equal(t, minRes, minExp)
	assert.Equal(t, maxRes, maxExp)

	assert.Equal(t, false, ok)

}
