package calculator

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Multi(a int, b int) int {
	return a * b
}

func Div(a int, b int) int {
	if b == 0 {
		panic("'b' can not be equals to zero")
	}

	return a / b
}
