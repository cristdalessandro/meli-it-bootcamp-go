package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageGrades(t *testing.T) {
	want := 6.0
	got := gradesAverage(6, 10, 4, 4, 6)

	assert.Equal(t, want, got)

}
