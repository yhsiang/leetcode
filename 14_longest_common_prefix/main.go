package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
