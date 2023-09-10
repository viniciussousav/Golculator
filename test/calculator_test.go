package test

import (
	"Golculator/src/calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculator(t *testing.T) {

	a, b := 5, 5

	t.Run("Sum", func(t *testing.T) {

		result := calculator.Sum(a, b)

		expected := 10

		assert.Equal(t, expected, result)
	})

	t.Run("Sub", func(t *testing.T) {

		result := calculator.Sub(a, b)

		expected := 0

		assert.Equal(t, expected, result)
	})

	t.Run("Multi", func(t *testing.T) {

		result := calculator.Multi(a, b)

		expected := 25

		assert.Equal(t, expected, result)
	})

	t.Run("Div", func(t *testing.T) {

		result := calculator.Div(a, b)

		expected := 1

		assert.Equal(t, expected, result)
	})
}
