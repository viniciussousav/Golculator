package main

import (
	"Golculator/src/calculator"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("____________________________________________________")
	fmt.Println("Welcome to first calculator in go by Vinicius Vilela")
	fmt.Println("____________________________________________________")

	run := true
	for run {
		printOptions()
		option := getInput(scanner)

		if !checkOption(option) {
			fmt.Println("Invalid option, try again")
			continue
		}

		fmt.Println("Type the first number")
		inputA := getInput(scanner)
		a, valid := checkNumberInput(inputA)
		if !valid {
			continue
		}

		fmt.Println("Type the second number")
		inputB := getInput(scanner)
		b, valid := checkNumberInput(inputB)
		if !valid {
			continue
		}

		var result string
		switch option {
		case "+":
			result = fmt.Sprintf("%d + %d = %d", a, b, calculator.Sum(a, b))
			break
		case "-":
			result = fmt.Sprintf("%d - %d = %d", a, b, calculator.Sub(a, b))
			break
		case "*":
			result = fmt.Sprintf("%d * %d = %d", a, b, calculator.Multi(a, b))
			break
		case "/":
			result = fmt.Sprintf("%d / %d = %d", a, b, calculator.Div(a, b))
			break
		default:
			result = fmt.Sprintf("An unexpected error occurred")
			break
		}

		fmt.Println("Completed: ", result)

		validOption := false
		for !validOption {
			fmt.Println("Do you want to try again? (y/N)")
			answer := strings.ToLower(getInput(scanner))
			validOption = answer == "y" || answer == "n"

			if !validOption {
				fmt.Println("Option not recognized. Please, try again!")
				continue
			}

			validOption = true
			run = answer == "y"
		}

		cleanConsole()

		if !run {
			fmt.Println("See you soon! Thank you for testing")
		}
	}
}

func getInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func printOptions() {
	fmt.Println("\nAvailable operations:")
	fmt.Println("'+' -> Sum")
	fmt.Println("'-' -> Sub")
	fmt.Println("'*' -> Mult")
	fmt.Println("'/' -> Div")
	fmt.Print("\nType an option: ")
}

func checkOption(input string) bool {
	options := []string{"+", "-", "/", "*"}
	return slices.Contains(options, input)
}

func checkNumberInput(input string) (number int, valid bool) {
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		valid = false
	} else {
		valid = true
	}
	return
}

func cleanConsole() {
	fmt.Println("\033[2J")
}
