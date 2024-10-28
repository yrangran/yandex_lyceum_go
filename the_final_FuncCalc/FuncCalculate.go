package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Calc evaluates a mathematical expression provided as a string and returns the result or an error if invalid.
func Calc(expression string) (float64, error) {
	nums := []float64{}
	ops := []rune{}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if unicode.IsSpace(ch) {
			continue
		}

		if unicode.IsDigit(ch) || ch == '.' {
			num, nextIdx, err := parseNumber(expression, i)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
			i = nextIdx - 1
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {
			if err := processUntilOpenParenthesis(&nums, &ops); err != nil {
				return 0, err
			}
		} else if isOperator(ch) {
			if err := processOperators(ch, &nums, &ops); err != nil {
				return 0, err
			}
		} else {
			return 0, errors.New("invalid character in expression")
		}
	}

	if err := finalizeCalculation(&nums, &ops); err != nil {
		return 0, err
	}

	if len(nums) != 1 {
		return 0, errors.New("invalid expression")
	}
	return nums[0], nil
}

// parseNumber parses a float64 number from the expression starting at index i.
func parseNumber(expression string, i int) (float64, int, error) {
	j := i
	for j < len(expression) && (unicode.IsDigit(rune(expression[j])) || expression[j] == '.') {
		j++
	}
	num, err := strconv.ParseFloat(expression[i:j], 64)
	if err != nil {
		return 0, j, errors.New("invalid number format")
	}
	return num, j, nil
}

// processUntilOpenParenthesis applies operators until an open parenthesis is found.
func processUntilOpenParenthesis(nums *[]float64, ops *[]rune) error {
	for len(*ops) > 0 && (*ops)[len(*ops)-1] != '(' {
		if err := applyOperator(nums, ops); err != nil {
			return err
		}
	}
	if len(*ops) == 0 {
		return errors.New("mismatched parentheses")
	}
	*ops = (*ops)[:len(*ops)-1] // Remove '('
	return nil
}

// processOperators handles operator precedence and applies operators if needed.
func processOperators(ch rune, nums *[]float64, ops *[]rune) error {
	for len(*ops) > 0 && precedence((*ops)[len(*ops)-1]) >= precedence(ch) {
		if err := applyOperator(nums, ops); err != nil {
			return err
		}
	}
	*ops = append(*ops, ch)
	return nil
}

// applyOperator applies the last operator to the last two numbers in the stack.
func applyOperator(nums *[]float64, ops *[]rune) error {
	if len(*nums) < 2 || len(*ops) == 0 {
		return errors.New("invalid expression")
	}
	b := (*nums)[len(*nums)-1]
	a := (*nums)[len(*nums)-2]
	*nums = (*nums)[:len(*nums)-2]
	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]

	var res float64
	switch op {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		if b == 0 {
			return errors.New("division by zero")
		}
		res = a / b
	}
	*nums = append(*nums, res)
	return nil
}

// finalizeCalculation applies all remaining operators in the stack.
func finalizeCalculation(nums *[]float64, ops *[]rune) error {
	for len(*ops) > 0 {
		if (*ops)[len(*ops)-1] == '(' {
			return errors.New("mismatched parentheses")
		}
		if err := applyOperator(nums, ops); err != nil {
			return err
		}
	}
	return nil
}

// precedence returns the precedence of the operator.
func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

// isOperator checks if the character is a valid operator.
func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func main() {
	expression := "3 + 5 * (3 * 7) / 2"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%f", result)
	}
}