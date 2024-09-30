package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	i := bufio.NewScanner(os.Stdin)
	i.Split(bufio.ScanLines)

	sp := 0
	entries := 0

	for i.Scan() {
		t := i.Text()
		views := strings.Split(t, " ")

		if len(views) == 1 {
			entries, _ = strconv.Atoi(views[0])
			continue
		}

		if (views[0] == "1" && views[1] == "1") || (views[0] == "1" && views[2] == "1") || (views[1] == "1" && views[2] == "1") {
			sp++
		}

		if entries == 1 {
			break
		}

		entries--
	}

	fmt.Println(sp)
}
