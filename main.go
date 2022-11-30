package main

import (
	"avarpuk_calc/calc"
	"avarpuk_calc/parser"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Consts
const colorClear = "\033[H\033[2J"

func main() {
	fmt.Print(colorClear)

	parsers := []parser.Parsers{
		&parser.ArabicParser{},
		&parser.RomanParser{},
	}

	calculators := calculator.NewDefaultCalculator()

	for {
		var (
			inputString, operation string
			num1, num2             float64
		)

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			inputString = strings.ToLower(scanner.Text())
		}

		switch inputString {
		case "clear":
			fmt.Print(colorClear)
		case "break":
			return
		default:
			for i := 0; i < len(parsers); i++ {
				if err := parsers[i].Parse(inputString, &num1, &operation, &num2); err == nil {
					break
				}
			}
		}
		calculators.Run(num1, operation, num2)
		fmt.Println(calculators.GetResult())
	}

}
