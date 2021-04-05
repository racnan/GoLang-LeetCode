//Runtime: 1292ms; Runtime Percentile: 5.14%
//Memory: 6.5MB; Memory Percentile: 19.13%

func lengthOfLongestSubstring(s string) int {
	str := ""
	max := 0
	match := false
	for i := 0; i < len(s); i++ {
		str += string(s[i])
		for j := i + 1; j < len(s); j++ {
			for _, ch := range str {
				if string(s[j]) == string(ch) {
					match = true
					break
				}
			}
			if match {
				break
			}
			str += string(s[j])
		}
		match = false
		if len(str) > max {
			max = len(str)
		}

		str = ""
	}
	return max
}
