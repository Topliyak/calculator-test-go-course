package main

import (
	"errors"
	"strconv"
	"strings"
)

const (
	actions = "+-*/"
)

func Calculate(expression string) (string, error) {
	left, right, action, err := SplitByAction(expression)

	if err != nil {
		return "", err
	}

	lv, err1 := GetNumber(left)
	rv, err2 := GetNumber(right)

	if err1 != nil {
		return "", err1
	}

	if err2 != nil {
		return "", err2
	}

	if lv < 1 || lv > 10 || rv < 1 || rv > 10 {
		return "", errors.New("Число должно быть между 1 и 10 включительно.")
	}

	if IsRomanic(left) != IsRomanic(right) {
		return "", errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}

	v := action.Execute(lv, rv)

	if IsRomanic(left) {
		return ToRomanic(v)
	}

	return strconv.FormatInt(int64(v), 10), nil
}

func SplitByAction(expression string) (left string, right string, action Action, e error) {
	act_index, err := GetActionIndex(expression)

	if err != nil {
		return "", "", nil, err
	}

	left = expression[:act_index]
	left = strings.TrimSpace(left)

	right = expression[act_index+1:]
	right = strings.TrimSpace(right)

	action = GetAction(expression[act_index])

	return left, right, action, nil
}

func GetActionIndex(expression string) (int, error) {
	index := -1

	for i, c := range expression {
		if strings.ContainsRune(actions, c) {
			if index == -1 {
				index = i
			} else {
				return -1, errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			}
		}
	}

	if index == -1 {
		return -1, errors.New("Вывод ошибки, так как строка не является математической операцией.")
	}

	return index, nil
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
