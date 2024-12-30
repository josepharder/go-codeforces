package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	isUpperCase := func(v byte) bool {
		return v >= 'A' && v <= 'Z'
	}

	toLower := func(v byte) byte {
		return v + ('a' - 'A')
	}

	sum := func(v byte) byte {
		if isUpperCase(v) {
			return (toLower(v))
		}

		return v
	}

	stringA := make([]byte, len(a))
	stringB := make([]byte, len(b))
	for i := range a {
		stringA = append(stringA, sum(a[i]))
		stringB = append(stringB, sum(b[i]))
	}

	a = string(stringA)
	b = string(stringB)

	result := 0
	if a > b {
		result = 1
	} else if a < b {
		result = -1
	}

	fmt.Println(result)
}
