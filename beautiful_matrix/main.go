package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	output := func(rowWithNumberOne, columnWithNumberOne int) float64 {
		minRowMoves := math.Abs(float64(3 - rowWithNumberOne))
		minColumnMoves := math.Abs(float64(3 - columnWithNumberOne))
		return minRowMoves + minColumnMoves
	}

	var minMoves *float64
	var i int

	for minMoves == nil {
		scanner.Scan()
		row := strings.Split(scanner.Text(), " ")
		i++

		for j, value := range row {
			if value == "1" {
				result := output(i, j+1)
				minMoves = &result
				break
			}
		}
	}

	fmt.Println(*minMoves)
}
