package day04_longest_common_prefix

func longestCommonPrefix(strs []string) string {
	minLenStr := strs[0]
	for _, str := range strs {
		if len(str) < len(minLenStr) {
			minLenStr = str
		}
	}
	for i := 0; i < len(minLenStr); i++ {
		for _, str := range strs {
			if str[i] != minLenStr[i] {
				return minLenStr[:i]
			}
		}
	}
	return minLenStr
}
