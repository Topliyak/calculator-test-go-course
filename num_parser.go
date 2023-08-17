package main

import (
	"strconv"
	"strings"
)

func IsNumber(expression string) bool {
	return IsDec(expression) || IsRomanic(expression)
}

func GetNumber(expression string) (int, error) {
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
