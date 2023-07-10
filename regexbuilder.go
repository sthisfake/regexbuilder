// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	keyword := "fev"
// 	condition := true

// 	// Dynamically build the regular expression pattern
// 	pattern := fmt.Sprintf(".*%s.*", regexp.QuoteMeta(keyword))

// 	if condition {
// 		pattern = fmt.Sprintf("^%s$", pattern)
// 	} else {
// 		pattern = fmt.Sprintf(".*%s.*", pattern)
// 	}

// 	fmt.Println(pattern)

// 	// Compile the regular expression
// 	regex, err := regexp.Compile(pattern)
// 	if err != nil {
// 		// Handle error if the pattern is invalid
// 		fmt.Printf("Error compiling regex pattern: %s\n", err.Error())
// 		return
// 	}

// 	// Perform matching or other operations with the compiled regex
// 	text := "This is an example text."
// 	if regex.MatchString(text) {
// 		fmt.Println("Text matches the pattern!")
// 	} else {
// 		fmt.Println("Text does not match the pattern.")
// 	}
// }

package regexbuilder

import "fmt"

type Condition struct {
	Number NumberCondition
	Text   TextCondition
}


type NumberCondition struct{}
type TextCondition struct{}


//Number Condition methods

func (NumberCondition) GreaterThan(value string) string {
	return fmt.Sprintf(`^[2-9]\d*%s`, value)
}


// Text Condition methods
func (TextCondition) ContainStatement(statement string) string {
	return fmt.Sprintf(`(?i)%s`, statement)
}

// BuildPattern generates a regular expression pattern based on the provided condition and value.
func BuildPattern(condition interface{}, value string) (string, error) {
	switch c := condition.(type) {
	case NumberCondition:
		return c.GreaterThan(value), nil
	case TextCondition:
		return c.ContainStatement(value), nil
	default:
		return "", fmt.Errorf("unsupported condition")
	}
}
