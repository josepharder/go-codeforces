package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	fmt.Println(getApprovedParticipants(n, k, scores))
}

func getApprovedParticipants(n, k int, scores []int) int {
	threshold := scores[k-1]
	condition := func(i int) bool {
		return scores[i] < threshold || scores[i] == 0
	}

	return binarySearch(n, condition)
}

func binarySearch(n int, f func(int) bool) int {
	min, max := 0, n

	for min < max {
		mid := int(uint(min+max) >> 1)
		if !f(mid) {
			min = mid + 1
		} else {
			max = mid
		}
	}

	return min
}
