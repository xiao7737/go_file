package dp_lcs_1143

import (
	"reflect"
	"testing"
)

func TestLongestCommonSubsequence3(t *testing.T) {
	var tests = []struct {
		inputOne string
		inputTwo string
		want     []string
	}{
		{"abcde", "abc", []string{"abc"}},
		{"abbc", "bb", []string{"bb"}},
		{"abca", "abeca", []string{"ab", "ca"}}, //有两个结果的情况
		{"abc", "d", []string{}},
		{"", "", []string{}},
	}
	for _, test := range tests {
		if got := LongestCommonSubsequence3(test.inputOne, test.inputTwo); !reflect.DeepEqual(got, test.want) {
			t.Errorf("LongestCommonSubsequence3(%q,%q)=%q", test.inputOne, test.inputTwo, got)
		}
	}
}
