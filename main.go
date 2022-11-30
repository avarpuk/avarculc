package main

import (
	calculator "avarpuk_calc/calc"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Consts
const colorClear = "\033[H\033[2J"

func main() {
	calculator := calculator.NewDefaultCalculator()

	for {
		var (
			inputString string
		)

		fmt.Print(colorClear)

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

		}
		fmt.Println(inputString)

	}

}
