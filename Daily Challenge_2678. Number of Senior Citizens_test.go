package leetcodego

import (
	"testing"
)
var testcase = []struct{
	name string
	details []string
	expected int
}{
	{"testcase1", []string{"7868190130M7522","5303914400F9211","9273338290F4010"}, 2},
	{"testcase2", []string{"1313579440F2036","2921522980M5644"}, 0},
}

func TestcountSeniors(t *testing.T){
	for _, tc := range testcase{
		t.Run(tc.name, func(t *testing.T){
			got := countSeniors(tc.details)
			if tc.expected != got{
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}

