package leetcodego

import(
	"testing"
)

func TestMinimumPushes(t *testing.T){
	var testCase = []struct{
		name string
		word string
		expected int
	}{
		{"testcase1", "abcde", 5},
		{"testcase2", "xyzxyzxyzxyz", 12},
		{"testcase3", "aabbccddeeffgghhiiiiii", 24},
	}

	for _, tc := range testCase{
		t.Run(tc.name, func(t *testing.T) {
			got := minimumPushes(tc.word)

			if tc.expected != got{
				t.Errorf("input = %s, expected = %d, but got = %d", tc.word, tc.expected, got)
			}
		})
	}
}