package dp_lcs_1143

import "testing"

func TestLongestCommonSubsequence3(t *testing.T) {
	var tests = []struct {
		inputOne string
		inputTwo string
		want     string
	}{
		{"abcde", "abc", "abc"},
		{"abbc", "bb", "bb"},
		{"abca", "abeca", "ab"}, //应该返回ab，ca
		{"abc", "d", ""},
		{"", "", ""},
	}
	for _, test := range tests {
		if got := LongestCommonSubsequence3(test.inputOne, test.inputTwo); got != test.want {
			t.Errorf("LongestCommonSubsequence3(%q,%q)=%q", test.inputOne, test.inputTwo, got)
		}
	}
}
