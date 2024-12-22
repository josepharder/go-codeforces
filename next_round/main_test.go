package main

import "testing"

func TestGetApprovedParticipants(t *testing.T) {
	tests := []struct {
		k        int
		scores   []int
		expected int
	}{
		{5, []int{10, 9, 8, 7, 7, 7, 5, 5}, 6},
		{2, []int{0, 0, 0, 0}, 0},
		{1, []int{1, 1, 1, 1, 1}, 5},
		{1, []int{10}, 1},
		{14, []int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 14},
		{3, []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}, 7},
	}

	for _, test := range tests {
		got := getApprovedParticipants(len(test.scores), test.k, test.scores)
		if got != test.expected {
			t.Errorf("k=%d, scores=%v => got %d, expected %d", test.k, test.scores, got, test.expected)
		}
	}
}
