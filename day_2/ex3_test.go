package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalaryCalculation(t *testing.T) {
	cataReal := salaryCalculation(180, 'A')
	catbReal := salaryCalculation(180, 'B')
	catcReal := salaryCalculation(180, 'C')

	cataExpected, catbExpected, catcExpected := 3000.0, 5400.0, 13500.0

	assert.Equal(t, cataReal, cataExpected)
	assert.Equal(t, catbReal, catbExpected)
	assert.Equal(t, catcReal, catcExpected)
}
