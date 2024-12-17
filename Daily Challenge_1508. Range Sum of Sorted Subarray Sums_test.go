package leetcodego

import (
	"testing"
)

func Test_rangeSum(t *testing.T) {
	var testcase = []struct {
		name     string
		nums     []int
		n        int
		left     int
		right    int
		expected int
	}{
		{"testcase1", []int{1, 2, 3, 4}, 4, 1, 5, 13},
		{"testcase2", []int{1, 2, 3, 4}, 4, 3, 4, 6},
		{"testcase3", []int{1, 2, 3, 4}, 4, 1, 10, 50},
	}
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			got := rangeSum(tc.nums, tc.n, tc.left, tc.right)
			if tc.expected != got {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
