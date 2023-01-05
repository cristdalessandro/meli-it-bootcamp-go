package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {

	got1 := calculateTaxes(30000.0)
	got2 := calculateTaxes(75000.0)
	got3 := calculateTaxes(205000.0)
	want1 := 0.0
	want2 := 12750.000000000002
	want3 := 55350.00000000001

	assert.Equal(t, got1, want1)
	assert.Equal(t, got2, want2)
	assert.Equal(t, got3, want3)
}
