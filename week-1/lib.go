package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	PARAMS_NUM = 3
	FIRST_NUM  = 0
	SECOND_NUM = 2
	OP         = 1
)

func getPriority(param string) int {
	switch param {
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	case "(":
		return -1
	case ")":
		return -2
	default:
		return 0
	}
}

func isOperator(op string) bool {
	switch op {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

func normalization(exp string) string {
	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(exp, " ")
	return strings.Trim(s, " ")
}

func getNumber(str string) (num float64, err error) {
	num, err = strconv.ParseFloat(str, 10)
	if err != nil {
		err = errors.New("Invalid input")
	}
	return
}

func cal(firstNum, secondNum float64, op string) (float64, error) {
	var result float64
	switch op {
	case "+":
		result = firstNum + secondNum
	case "-":
		result = firstNum - secondNum
	case "*":
		result = firstNum * secondNum
	case "/":
		if secondNum == 0 {
			return result, errors.New("Cannot divide by zero")
		}
		result = firstNum / secondNum
	default:
		return result, errors.New("Invalid operation!! Must be one of +, -, *, /")
	}
	return result, nil
}

// Evaluate calculate result from user expression
func Evaluate(exp *string) (float64, error) {
	*exp = normalization(*exp)
	params := strings.Split(*exp, " ")
	stack := make([]string, 0)
	postfix := make([]string, 0)

	// Convert to prefix
	for _, param := range params {
		switch prio := getPriority(param); prio {
		case 0: // if number add to postfix
			postfix = append(postfix, param)
		case 1, 2: // if operator and add to stack
			//compare to stack, if stack's op priority higher pop stack op
			if len(stack) > 0 && getPriority(stack[0]) >= prio {
				postfix = append(postfix, stack[0])
				stack = stack[1:]
			}
			stack = append([]string{param}, stack...)
		case -1: // if "(" push to stack
			stack = append([]string{param}, stack...)
		case -2: // if ")" pop all stack ulti meet "("
			for len(stack) > 0 && stack[0] != "(" {
				postfix = append(postfix, stack[0])
				stack = stack[1:]
			}

			if len(stack) > 0 {
				stack = stack[1:]
			} else {
				return 0, errors.New("Syntax error")
			}
		}
	}

	for _, op := range stack {
		postfix = append(postfix, op)
	}
	// end convert

	// Calculate on postfix
	calStack := make([]float64, 0)
	for _, param := range postfix {
		if isOperator(param) {
			if len(calStack) < 2 {
				return 0, errors.New("Syntax error")
			}

			first := calStack[0]
			second := calStack[1]
			result, err := cal(first, second, param)
			if err != nil {
				return result, err
			}

			calStack[1] = result
			calStack = calStack[1:]
		} else {
			num, err := getNumber(param)
			if err != nil {
				return 0, err
			}
			calStack = append([]float64{num}, calStack...)
		}
	}

	if len(calStack) > 1 {
		return 0, errors.New("Syntax error")
	}

	return calStack[0], nil
}
