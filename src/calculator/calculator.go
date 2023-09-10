package calculator

// Sum Returns the sum of two numbers
func Sum(a int, b int) int {
	return a + b
}

// Sub Returns the subtraction of two numbers
func Sub(a int, b int) int {
	return a - b
}

// Multi Returns the multiplication of two numbers
func Multi(a int, b int) int {
	return a * b
}

// Div Returns the division of two numbers
func Div(a int, b int) int {
	if b == 0 {
		panic("'b' can not be equals to zero")
	}

	return a / b
}
