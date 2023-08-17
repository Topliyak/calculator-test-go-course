package main

type Sum struct{}

func (Sum) Execute(left, right int) int {
	return left + right
}
