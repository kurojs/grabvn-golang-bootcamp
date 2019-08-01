package main

import (
	"errors"
	"strconv"
	"strings"
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

func isValidOperator(op string) bool {
	switch op {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
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
	params := strings.Fields(*exp)
	stack := make([]string, 0)
	postfix := make([]string, 0)

	// Convert to prefix
	for _, param := range params {
		switch prio := getPriority(param); prio {
		case 0: // if number add to postfix
			postfix = append(postfix, param)
		case 1, 2: // if operator and add to stack
			//compare to stack, if stack's op priority higher pop stack op
			if len(stack) > 0 && getPriority(stack[len(stack)-1]) >= prio {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, param)
		case -1: // if "(" push to stack
			stack = append(stack, param)
		case -2: // if ")" pop all stack ulti meet "("
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				return 0, errors.New("Syntax error")
			}
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	// end convert

	// Calculate on postfix
	calStack := make([]float64, 0)
	for _, param := range postfix {
		if isValidOperator(param) {
			if len(calStack) < 2 {
				return 0, errors.New("Syntax error")
			}

			first := calStack[len(calStack)-2]
			second := calStack[len(calStack)-1]
			calStack = calStack[:len(calStack)-2]

			result, err := cal(first, second, param)
			if err != nil {
				return result, err
			}

			calStack = append(calStack, result)
		} else {
			num, err := strconv.ParseFloat(param, 10)
			if err != nil {
				return 0, err
			}
			calStack = append(calStack, num)
		}
	}

	if len(calStack) > 1 {
		return 0, errors.New("Syntax error")
	}

	return calStack[len(calStack)-1], nil
}
