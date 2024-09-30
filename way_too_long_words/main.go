package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	i := bufio.NewScanner(os.Stdin)

	for i.Scan() {
		t := i.Text()

		if _, err := strconv.Atoi(t); err == nil {
			continue
		}

		if len(t) <= 10 {
			fmt.Println(t)
		} else {
			a := t[:1]
			rest := t[1 : len(t)-1]
			b := t[len(t)-1:]
			fmt.Printf("%s%d%s\n", a, len(rest), b)
		}
	}
}
