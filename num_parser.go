package main

import (
	"errors"
	"strconv"
	"strings"
)

func IsNumber(expression string) bool {
	return IsDec(expression) || IsRomanic(expression)
}

func GetNumber(expression string) (int, error) {
	if !IsNumber(expression) {
		return 0, errors.New("Принимаются только целые числа.")
	}

	if IsRomanic(expression) {
		return ParseRomanic(expression)
	}

	return ParseDec(expression)
}

func IsDec(expression string) bool {
	return ContainsOnly(expression, "1234567890")
}

func ParseDec(expression string) (int, error) {
	trimmed := strings.TrimSpace(expression)
	return strconv.Atoi(trimmed)
}

func IsRomanic(expression string) bool {
	return ContainsOnly(expression, "IVXLCDM")
}

func ContainsOnly(expression, chars string) bool {
	for _, c := range expression {
		if !strings.ContainsRune(chars, c) {
			return false
		}
	}

	return true
}

func ParseRomanic(expression string) (int, error) {
	rome_dec_pairs := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	trimmed := strings.TrimSpace(expression)
	prev := -1
	res := 0

	for _, rome := range trimmed {
		curr := rome_dec_pairs[rome]

		res += curr

		if prev < curr && prev != -1 {
			res -= 2 * prev
		}

		prev = curr
	}

	return res, nil
}

func ToRomanic(n int) (string, error) {
	if n < 0 {
		return "", errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
	}

	if n == 0 {
		return "", errors.New("Вывод ошибки, так как в римской системе нет нуля.")
	}

	dec_and_romanic_pairs := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "XI",
		10:  "X",
		20:  "XX",
		30:  "XXX",
		40:  "XL",
		50:  "L",
		60:  "LX",
		70:  "LXX",
		80:  "LXXX",
		90:  "XC",
		100: "C",
	}

	r := ""
	k := 10

	for n > 0 {
		d := n % k
		r = dec_and_romanic_pairs[d] + r
		n -= d
		k *= 10
	}

	return r, nil
}
