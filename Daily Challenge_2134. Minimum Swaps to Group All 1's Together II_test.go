package leetcodego

import (
	"testing"
)

var testcase2 = []struct {
	name     string
	nums     []int
	expected int
}{
	{"testcase1", []int{0, 1, 0, 1, 1, 0, 0}, 1},
	{"testcase2", []int{0, 1, 1, 1, 0, 0, 1, 1, 0}, 2},
	{"testcase3", []int{1, 1, 0, 0, 1}, 0},
}

func Test_minSwaps(t *testing.T) {
	for _, tc := range testcase2 {
		t.Run(tc.name, func(t *testing.T) {
			got := minSwaps(tc.nums)

			if tc.expected != got {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
