package leetcodego

import(
	"testing"
)

func TestKthDistinct(t *testing.T){
	var testCase = []struct{
		name string
		arr []string
		k int
		expected string
	}{
		{"testcase 1", []string{"d","b","c","b","c","a"}, 2, "a"},
		{"testcase 2", []string{"aaa","aa","a"}, 1, "aaa"},
		{"testcase 3", []string{"a","b","a"}, 3, ""},
	}

	for _, tc := range testCase{
		t.Run(tc.name, func(t *testing.T){
			got := kthDistinct(tc.arr, tc.k)

			if tc.expected != got{
				t.Errorf("With input arr = %v, and k = %d, expected %s, but got %s", tc.arr, tc.k, tc.expected, got)
			}
		})
	}
}