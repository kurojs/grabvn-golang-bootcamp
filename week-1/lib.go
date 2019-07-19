package main

import (
	"errors"
	"fmt"
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

// func eval(in string) res int, err error {
// 	stack := strings.Split(in, " ")
// 	isOperater := false
// 	for param, i := range stack {
// 		val, err := strconv.ParseFloat()
// 	}

// 	return
// }

func removeDupSpace(exp string) string {
	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(exp, " ")
	return strings.Trim(s, " ")
}

func getNumber(str string) (num float64, err error) {
	num, err = strconv.ParseFloat(str, 10)
	if err != nil {
		err = errors.New("Invalid input!! Must be a number")
	}
	return
}

// Evaluate calculate result from user expression
func Evaluate(exp *string) (result float64, err error) {
	*exp = removeDupSpace(*exp)
	params := strings.Split(*exp, " ")
	if len(params) != PARAMS_NUM {
		err = errors.New("Invalid params number")
	}

	firstNum, err := getNumber(params[FIRST_NUM])
	if err != nil {
		return
	}
	secondNum, err := getNumber(params[SECOND_NUM])
	if err != nil {
		return
	}

	switch op := params[OP]; op {
	case "+":
		result = firstNum + secondNum
	case "-":
		result = firstNum - secondNum
	case "*":
		result = firstNum * secondNum
	case "/":
		if secondNum == 0 {
			err = errors.New("Cannot divide by zero")
			return
		}
		result = firstNum / secondNum
	default:
		fmt.Println(params[OP])
		err = errors.New("Invalid operation!! Must be one of +, -, *, /")
		return
	}
	return
}
