package parser

import "fmt"

type RomanParser struct{}

func (p *RomanParser) Parse(inputString string, num1 *float64, operation *string, num2 *float64) error {
	if _, err := fmt.Sscanf(inputString, "%s %s %s", num1, operation, num2); err != nil {
		return err
	}

	// logic

	return nil
}
