// Runtime: 788ms; Runtime Percentile: 11.05%
// Memory: 3.5MB; Memory Percentile:  37.95%

func longestPalindrome(s string) string {
	maxSt := ""
	maxlen := 0
	for i := range s {
		fmt.Println("i:", i)
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if isPalindrome(i, j, &s) {
					if maxlen < max(i, j) {
						maxlen = max(i, j)
						maxSt = s[i : j+1]
					}
				}
			}
		}
	}
	return string(maxSt)
}

func isPalindrome(a int, b int, st *string) bool {
	for a <= b {
		if (*st)[a] != (*st)[b] {
			return false
		}
		a++
		b--
	}
	return true
}
func max(a int, b int) int {
	if a == b {
		return 1
	} else {
		return 1 + b - a
	}
}