package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")

	n, _ := strconv.ParseFloat(input[0], 64)
	m, _ := strconv.ParseFloat(input[1], 64)
	a, _ := strconv.ParseFloat(input[2], 64)

	minLength := math.Ceil(n / a)
	minWidth := math.Ceil(m / a)
	minSquareFlags := int(minLength * minWidth)

	fmt.Println(minSquareFlags)
}
