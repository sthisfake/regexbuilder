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

import (
	"fmt"
	"reflect"
)

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

func (NumberCondition) LessThan(value string) string {
	return fmt.Sprintf(`^(1[0-9]|2[0-9])%s`, value)
}

func (NumberCondition) EvenNumber(value string) string {
	return fmt.Sprintf(`^[24680]\d*%s`, value)
}

func (NumberCondition) OddNumber(value string) string {
	return fmt.Sprintf(`^[13579]\d*%s`, value)
}

// Text Condition methods

func (TextCondition) ContainStatement(statement string) string {
	return fmt.Sprintf(`(?i).*%s.*`, statement)
}

func (TextCondition) WithoutStatement(statement string) string {
	return fmt.Sprintf(`(?i)^(?!.*%s).*`, statement)
}

func (TextCondition) GreaterCharacterSizeThan(statement string) string {
	return fmt.Sprintf(`^(.{%s})$`, statement)
}

func (TextCondition) LessCharacterSizeThan(statement string) string {
	return fmt.Sprintf(`^(.{0,%s})$`, statement)
}


// BuildPattern generates a regular expression pattern based on the provided condition and value.
func BuildPattern(condition interface{}, value string) (string, error) {
	conditionValue := reflect.ValueOf(condition)
	method := conditionValue.MethodByName("BuildPattern")
	if method.IsValid() {
		result := method.Call([]reflect.Value{reflect.ValueOf(value)})
		if len(result) == 2 {
			pattern, ok := result[0].Interface().(string)
			if !ok {
				return "", fmt.Errorf("invalid pattern type")
			}
			err, ok := result[1].Interface().(error)
			if !ok {
				return "", fmt.Errorf("invalid error type")
			}
			if err != nil {
				return "", err
			}
			return pattern, nil
		}
	}

	return "", fmt.Errorf("unsupported condition")
}
