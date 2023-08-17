package main

import "strings"

const (
	actions = "+-*/"
)

func Calculate(expression string) int {
	expression = DeleteBrackets(expression)

	if IsNumber(expression) {
		n, _ := GetNumber(expression)
		return n
	}

	left, right, action := SplitByAction(expression)

	lv := Calculate(left)
	rv := Calculate(right)

	return action.Execute(lv, rv)
}

func SplitByAction(expression string) (left string, right string, action Action) {
	act_index := GetActionIndex(expression)

	left = expression[:act_index]
	right = expression[act_index+1:]
	action = GetAction(expression[act_index])

	return left, right, action
}

func DefineActionChar(expression string) rune {
	for _, a := range actions {
		if IsActionThere(expression, a) {
			return a
		}
	}

	return ' '
}

func GetActionIndex(expression string) int {
	action := DefineActionChar(expression)
	opened_brackets := 0

	for i, c := range expression {
		if c == '(' {
			opened_brackets++
		}

		if c == ')' {
			opened_brackets--
		}

		if opened_brackets == 0 && c == action {
			return i
		}
	}

	return -1
}

func IsActionThere(expression string, action rune) bool {
	opened_brackets := 0

	for _, c := range expression {
		if c == '(' {
			opened_brackets++
		}

		if c == ')' {
			opened_brackets--
		}

		if opened_brackets == 0 && c == action {
			return true
		}
	}

	return false
}

func DeleteBrackets(expression string) string {
	expression = strings.TrimSpace(expression)

	for InBrackets(expression) {
		end := len(expression) - 1
		expression = expression[1:end]
	}

	return expression
}

func InBrackets(expression string) bool {
	end := len(expression) - 1
	cutted := expression[:end]
	brackets_opened := 0

	for _, c := range cutted {
		if c == '(' {
			brackets_opened++
		}

		if c == ')' {
			brackets_opened--
		}
	}

	return brackets_opened == 1
}

func GetAction(c byte) Action {
	switch c {
	case '+':
		return Sum{}
	case '-':
		return Difference{}
	case '*':
		return Multiplication{}
	case '/':
		return Division{}
	default:
		return nil
	}
}

type Action interface {
	Execute(left, right int) int
}
