package parser

import (
	"fmt"
)

type ArabicParser struct{}

func (p *ArabicParser) Parse(inputString string, num1 *float64, operation *string, num2 *float64) error {
	if _, err := fmt.Sscanf(inputString, "%v %s %v", num1, operation, num2); err != nil {
		return err
	}

	// logic

	return nil
}
