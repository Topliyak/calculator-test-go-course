package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	exp, _ := reader.ReadString('\n')
	res, err := Calculate(exp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// reader.ReadString('\n')
}
