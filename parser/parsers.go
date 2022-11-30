package parser

type Parsers interface {
	Parse(strToParse string, num1 *float64, operation *string, num2 *float64) error
}
