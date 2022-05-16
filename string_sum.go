package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func checkSymbols (input string) bool {

	allowed := "1234567890+- "

	for _, val := range(input) {

		if ok := strings.Contains(allowed, string(val)); !ok {
			return false
		}

	}

	return true
}

func parseString(input string) (num1 int, num2 int, err error) {

	re := regexp.MustCompile(`.?\d+`)
	res := re.FindAllString(input, -1)

	if len(res) != 2 {
		err = fmt.Errorf("bad input %w", errorNotTwoOperands)
		return		
	}

	num1, err1 := strconv.Atoi(res[0])
	if err1 != nil {
		return 0, 0, fmt.Errorf("something wrong, %w", err1)
	}

	num2, err2 := strconv.Atoi(res[1])
	if err2 != nil {
		return 0, 0, fmt.Errorf("something wrong, %w", err1)
	}

	return
	
}

func StringSum(input string) (output string, err error) {

	if ok := checkSymbols(input); !ok {
		err = fmt.Errorf("bad input. %w", strconv.ErrSyntax)
		return 
	}

	clearInput := strings.ReplaceAll(input, " ", "")

	if clearInput == ""{
		err = fmt.Errorf("bad input. %w",  errorEmptyInput)
		return
	}

	num1, num2, err := parseString(clearInput)
	if err != nil {
		return
	}

	result := num1 + num2

	output = strconv.Itoa(result)

	return
}
