package main

type Multiplication struct{}

func (Multiplication) Execute(left, right int) int {
	return left * right
}
