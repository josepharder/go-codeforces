package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	operations := 0
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &operations)

	x := 0
	for operations > 0 {
		scanner.Scan()
		statement := scanner.Text()

		if string(statement[1]) == "+" {
			x++
		} else {
			x--
		}

		operations--
	}

	fmt.Println(x)

}
