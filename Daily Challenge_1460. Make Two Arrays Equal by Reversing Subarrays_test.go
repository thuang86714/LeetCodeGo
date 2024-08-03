package leetcodego

import (
	"testing"
)

var testcase = []struct {
	name     string
	target   []int
	arr      []int
	expected bool
}{
	{"testcase1", []int{1, 2, 3, 4}, []int{2, 4, 1, 3}, true},
	{"testcase2", []int{7}, []int{7}, true},
	{"testcase3", []int{3, 7, 9}, []int{3, 7, 11}, false},
}

func Test_canBeEqual(t *testing.T) {
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			got := canBeEqual(tc.target, tc.arr)
			if tc.expected != got {
				t.Errorf("expected %t, got %t", tc.expected, got)
			}
		})
	}
}
