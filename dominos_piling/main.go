package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	squares := m * n
	fmt.Println(squares / 2)
}
