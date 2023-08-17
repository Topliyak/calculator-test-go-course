package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	exp, _ := reader.ReadString('\n')
	res := Calculate(exp)
	fmt.Println(res)

	reader.ReadString('\n')
}
