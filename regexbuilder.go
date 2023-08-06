package regexbuilder

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NumberConditionType struct{}
type TextConditionType struct{}

var NumberCondition NumberConditionType
var TextCondition TextConditionType




type Condition struct {
	Number NumberConditionType
	Text   TextConditionType
}



type NineSituations struct {
	allNineFlag bool
	lastDigitNine bool
	oneToLastDigitNine bool
	lastDigitZero bool

}


//Number Condition methods

func (NumberConditionType) GreaterThan(value string) (string  , error) {

	var regexPattern string
	var eachCondition []string

	cleanStringNumber := strings.ReplaceAll(value, " ", "")
	numberLen := len(cleanStringNumber)

	_, err := strconv.Atoi(cleanStringNumber)
	if err != nil {
		return "" , err
	}

	if numberLen == 1 {
		digitInt , err := strconv.Atoi(string(cleanStringNumber[0])) 
		if err != nil {
			return "" , err
		}
		stringDigit := strconv.Itoa(digitInt + 1)
		fmt.Println(stringDigit)
		regexPattern = fmt.Sprintf("^([%s-9]|[1-9]\\d{1,})$", stringDigit  )
		fmt.Println(regexPattern)
		return regexPattern , nil
	}

	digitInt , err := strconv.Atoi(string(cleanStringNumber[0])) 
	if err != nil {
		return "" , err
	}
	if digitInt == 0 {
		return "" , errors.New("first digit zero error")
	}


	situations := numberType(cleanStringNumber)

	if situations.allNineFlag  {
		numberLenString := strconv.Itoa(numberLen)
		regexPattern = fmt.Sprintf("^([1-9]\\d{%s,})$", numberLenString  )
		return regexPattern , nil
	}

	digits:= []byte(cleanStringNumber)
	var cacheString string
	for i:=0 ; i< len(digits) ; i++ {

		// for the first itteration :
		if i==0 {
			// check if the first digit is 9 
			if number , _ := strconv.Atoi(string(digits[i])) ;  number != 9{
				//if so , do the operation for the biggest digit 
				temp  ,_ := strconv.Atoi(string(digits[i])) 
				newNumber   := strconv.Itoa(temp + 1 )
				bigPattern := fmt.Sprintf("[%s-9]", newNumber  )
				sth := strings.Repeat("[0-9]" , len(digits) - 1 )
				bigPattern = bigPattern + sth
				eachCondition = append(eachCondition, bigPattern)
				thisNumber := strconv.Itoa(temp )
				cacheString = fmt.Sprintf("[%s]" , thisNumber )
			}else{
				thisNumber := strconv.Itoa(number )
				cacheString = cacheString + fmt.Sprintf("[%s]" , thisNumber )
				;
			} 
		}else {
			if number , _ := strconv.Atoi(string(digits[i])) ;  number != 9{
				temp  ,_ := strconv.Atoi(string(digits[i])) 
				newNumber   := strconv.Itoa(temp + 1 )
				bigPattern := fmt.Sprintf("[%s-9]", newNumber  )
				sth := strings.Repeat("[0-9]" , len(digits) - (i+1) )
				finalPattern := cacheString + bigPattern + sth
				eachCondition = append(eachCondition, finalPattern)
				thisNumber := strconv.Itoa(temp )
				cacheString = cacheString + fmt.Sprintf("[%s]" , thisNumber )
			}else{
				thisNumber := strconv.Itoa(number )
				cacheString = cacheString + fmt.Sprintf("[%s]" , thisNumber )
			}
		}
	}
	
	for index ,item :=  range eachCondition {
		if index == 0 {
			regexPattern = item + "|"
		}
		regexPattern = regexPattern + item + "|"
	}

	numberLenString := strconv.Itoa(numberLen)
	regexPattern = "^(" + regexPattern + "[1-9]\\d" + "{" + numberLenString +  ",})$"

	return regexPattern , err
}

func (NumberConditionType) LesserThan(value string) (string, error) {
	var regexPattern string
	var eachCondition []string

	cleanStringNumber := strings.ReplaceAll(value, " ", "")
	numberLen := len(cleanStringNumber)

	_, err := strconv.Atoi(cleanStringNumber)
	if err != nil {
		return "", err
	}

	if numberLen == 1 {
		digitInt, err := strconv.Atoi(string(cleanStringNumber[0]))
		if err != nil {
			return "", err
		}
		if digitInt == 0 {
		
			regexPattern = "^$"
			return regexPattern, nil
		}
		stringDigit := strconv.Itoa(digitInt - 1)
		regexPattern = fmt.Sprintf("^([0-%s])$", stringDigit)
		return regexPattern, nil
	}

	digitInt, err := strconv.Atoi(string(cleanStringNumber[0]))
	if err != nil {
		return "", err
	}
	if digitInt == 0 {
		return "", errors.New("first digit zero error")
	}

	digits := []byte(cleanStringNumber)
	var cacheString string
	for i := 0; i < len(digits); i++ {
		// for the first iteration:
		if i == 0 {
			// check if the first digit is 1
			if number, _ := strconv.Atoi(string(digits[i])); number != 0 {
				// if not, do the operation for the smallest digit
				temp, _ := strconv.Atoi(string(digits[i]))
				newNumber := strconv.Itoa(temp - 1)
				smallPattern := fmt.Sprintf("[0-%s]", newNumber)
				sth := strings.Repeat("[0-9]", len(digits)-1)
				smallPattern = smallPattern + sth
				eachCondition = append(eachCondition, smallPattern)
				thisNumber := strconv.Itoa(temp)
				cacheString = fmt.Sprintf("[%s]", thisNumber)
			} else {
				thisNumber := strconv.Itoa(number)
				cacheString = cacheString + fmt.Sprintf("[%s]", thisNumber)
			}
		} else {
			if number, _ := strconv.Atoi(string(digits[i])); number != 0 {
				temp, _ := strconv.Atoi(string(digits[i]))
				newNumber := strconv.Itoa(temp - 1)
				smallPattern := fmt.Sprintf("[0-%s]", newNumber)
				sth := strings.Repeat("[0-9]", len(digits)-(i+1))
				finalPattern := cacheString + smallPattern + sth
				eachCondition = append(eachCondition, finalPattern)
				thisNumber := strconv.Itoa(temp)
				cacheString = cacheString + fmt.Sprintf("[%s]", thisNumber)
			} else {
				thisNumber := strconv.Itoa(number)
				cacheString = cacheString + fmt.Sprintf("[%s]", thisNumber)
			}
		}
	}

	for index, item := range eachCondition {
		if index == 0 {
			regexPattern = item + "|"
		}
		regexPattern = regexPattern + item + "|"
	}

	regexPattern = "^([0-9]|" + regexPattern + "[0-9]\\d{" + strconv.Itoa(numberLen-2) + "})$"

	return regexPattern, nil
}


func (NumberConditionType) LessThan(value string) string {
	return fmt.Sprintf(`^(1[0-9]|2[0-9])%s`, value)
}

func (NumberConditionType) EvenNumber(value string) string {
	return fmt.Sprintf(`^[24680]\d*%s`, value)
}

func (NumberConditionType) OddNumber(value string) string {
	return fmt.Sprintf(`^[13579]\d*%s`, value)
}

// Text Condition methods

func (TextConditionType) ContainStatement(statement string) string {
	return fmt.Sprintf(`.*%s.*`, statement)
}

func (TextConditionType) WithoutStatement(statement string) string {
	return fmt.Sprintf(`^(?!.*%s).*`, statement)
}

func (TextConditionType) GreaterCharacterSizeThan(statement string) string {
	return fmt.Sprintf("^(.{%s,})$", statement)
}

func (TextConditionType) LessCharacterSizeThan(statement string) string {
	return fmt.Sprintf(`^(.{0,%s})$`, statement)
}


// BuildPattern generates a regular expression pattern based on the provided condition and value.
func BuildPattern(condition interface{}, value string) (string, error) {

	conditionValue := reflect.ValueOf(condition)
	conditionResult  := conditionValue.Call([]reflect.Value{reflect.ValueOf(value)})
	
	if len(conditionResult) <= 2 {
		conditionPattern, ok := conditionResult[0].Interface().(string)
		if !ok {
			return "", fmt.Errorf("invalid pattern type")
		}

		return conditionPattern, nil
	}

	return "", fmt.Errorf("unexpected number of return values")
}

func numberType(numberString string) NineSituations {

	result := NineSituations{
		allNineFlag: true,
		lastDigitNine: true,
		oneToLastDigitNine: true,
		lastDigitZero: true,
	}

	something := []byte(numberString)
	lastDigit , _ := strconv.Atoi(string(something[len(something) -1 ])) 
	oneToLastDigit , _ := strconv.Atoi(string(something[len(something) -2 ])) 
	if lastDigit != 9 {
		result.lastDigitNine = false
		result.allNineFlag = false
	}
	if lastDigit != 0 {
		result.lastDigitZero = false
	}
	if oneToLastDigit != 9 {
		result.oneToLastDigitNine = false
		result.allNineFlag = false
	}

	if result.allNineFlag  {
		for _, digit := range []byte(numberString) {
			intDigit , _ := strconv.Atoi(string(digit)) 
			if intDigit != 9 {
				result.allNineFlag = false
			}
		}
	}

	return result
}