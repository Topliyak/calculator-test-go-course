package main

type Difference struct{}

func (Difference) Execute(left, right int) int {
	return left - right
}
