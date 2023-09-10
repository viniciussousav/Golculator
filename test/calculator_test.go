package test

import (
	"Golculator/src/calculator"
	"testing"
)

func TestCalculator(t *testing.T) {

	a, b := 5, 5

	t.Run("Sum", func(t *testing.T) {

		result := calculator.Sum(a, b)

		expected := 10
		if result != expected {
			t.Errorf("Expected %d instead of %d as result", expected, result)
		}
	})

	t.Run("Sub", func(t *testing.T) {

		result := calculator.Sub(a, b)

		expected := 0
		if result != expected {
			t.Errorf("Expected %d instead of %d as result", expected, result)
		}
	})

	t.Run("Multi", func(t *testing.T) {

		result := calculator.Multi(a, b)

		expected := 25
		if result != expected {
			t.Errorf("Expected %d instead of %d as result", expected, result)
		}
	})

	t.Run("Div", func(t *testing.T) {

		result := calculator.Div(a, b)

		expected := 1
		if result != expected {
			t.Errorf("Expected %d instead of %d as result", expected, result)
		}
	})
}
