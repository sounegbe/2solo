package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hist := NewHistory()

	fmt.Println("=================================")
	fmt.Println("  Advanced CLI Calculator")
	fmt.Println("=================================")
	for {
		displayMenu()
		fmt.Print("\nEnter your choice: ")

		if !scanner.Scan() {
			return
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			result := performBinaryOp(scanner, "Addition", Add)
			if result != nil {
				hist.Add(fmt.Sprintf("Addition: %.2f", *result), *result)
			}
		case "2":
			result := performBinaryOp(scanner, "Subtraction", Subtract)
			if result != nil {
				hist.Add(fmt.Sprintf("Subtraction: %.2f", *result), *result)
			}
		case "3":
			result := performBinaryOp(scanner, "Multiplication", Multiply)
			if result != nil {
				hist.Add(fmt.Sprintf("Multiplication: %.2f", *result), *result)
			}
		case "4":
			result := performDivision(scanner)
			if result != nil {
				hist.Add(fmt.Sprintf("Division: %.2f", *result), *result)
			}
		case "5":
			result := performPower(scanner)
			if result != nil {
				hist.Add(fmt.Sprintf("Power: %.2f", *result), *result)
			}
		case "6":
			result := performModulus(scanner)
			if result != nil {
				hist.Add(fmt.Sprintf("Modulus: %.2f", *result), *result)
			}
		case "7":
			result := performSqrt(scanner)
			if result != nil {
				hist.Add(fmt.Sprintf("Square Root: %.2f", *result), *result)
			}
		case "8":
			result := performAverage(scanner)
			if result != nil {
				hist.Add(fmt.Sprintf("Average: %.2f", *result), *result)
			}
		case "9":
			hist.Display()
		case "10":
			fmt.Println("\nThank you for using the calculator. Goodbye!")
			return
		default:
			fmt.Println("\n Invalid choice. Please select 1-10.")
		}
	}
}

func displayMenu() {
	fmt.Println("\n---------------------------------")
	fmt.Println("MENU:")
	fmt.Println("1.  Add")
	fmt.Println("2.  Subtract")
	fmt.Println("3.  Multiply")
	fmt.Println("4.  Divide")
	fmt.Println("5.  Power (x^y)")
	fmt.Println("6.  Modulus (x%y)")
	fmt.Println("7.  Square Root")
	fmt.Println("8.  Average")
	fmt.Println("9.  Show History")
	fmt.Println("10. Exit")
	fmt.Println("---------------------------------")
}

func getNumber(scanner *bufio.Scanner, prompt string) (*float64, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read input")
	}

	input := strings.TrimSpace(scanner.Text())
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid number: %s", input)
	}

	return &num, nil

}

func performBinaryOp(scanner *bufio.Scanner, opName string, op func(float64, float64) float64) *float64 {
	fmt.Printf("\n--- %s ---\n", opName)

	num1, err := getNumber(scanner, "Enter first number: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	num2, err := getNumber(scanner, "Enter second number: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	result := op(*num1, *num2)
	fmt.Printf(" Result: %.2f\n", result)
	return &result
}

func performDivision(scanner *bufio.Scanner) *float64 {
	fmt.Println("\n--- Division ---")

	num1, err := getNumber(scanner, "Enter dividend: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	num2, err := getNumber(scanner, "Enter divisor: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	result, err := Divide(*num1, *num2)
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	fmt.Printf("Result: %.2f\n", result)
	return &result
}

func performPower(scanner *bufio.Scanner) *float64 {
	fmt.Println("\n--- Power (x^y) ---")

	base, err := getNumber(scanner, "Enter base: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	exp, err := getNumber(scanner, "Enter exponent: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	result := Power(*base, *exp)
	fmt.Printf(" Result: %.2f\n", result)
	return &result
}

func performModulus(scanner *bufio.Scanner) *float64 {
	fmt.Println("\n--- Modulus ---")

	num1, err := getNumber(scanner, "Enter first number: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	num2, err := getNumber(scanner, "Enter second number: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	result, err := Modulus(*num1, *num2)
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	fmt.Printf(" Result: %.2f\n", result)
	return &result
}

func performSqrt(scanner *bufio.Scanner) *float64 {
	fmt.Println("\n--- Square Root ---")

	num, err := getNumber(scanner, "Enter number: ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	result, err := SquareRoot(*num)
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	fmt.Printf(" Result: %.2f\n", result)
	return &result
}

func performAverage(scanner *bufio.Scanner) *float64 {
	fmt.Println("\n--- Average ---")

	countStr, err := getNumber(scanner, "How many numbers? ")
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	count := int(*countStr)
	if count <= 0 {
		fmt.Println(" Error: count must be positive")
		return nil
	}

	numbers := make([]float64, 0, count)
	for i := 0; i < count; i++ {
		num, err := getNumber(scanner, fmt.Sprintf("Enter number %d: ", i+1))
		if err != nil {
			fmt.Printf(" Error: %v\n", err)
			return nil
		}
		numbers = append(numbers, *num)
	}

	result, err := Average(numbers)
	if err != nil {
		fmt.Printf(" Error: %v\n", err)
		return nil
	}

	fmt.Printf(" Result: %.2f\n", result)
	return &result
}
