package main

type Division struct{}

func (Division) Execute(left, right int) int {
	return left / right
}
