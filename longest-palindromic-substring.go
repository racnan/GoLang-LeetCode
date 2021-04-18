// Runtime: 4ms; Runtime Percentile: 96.07%
// Memory: 2.7MB; Memory Percentile: 46.74%

func longestPalindrome(s string) string {
	var a, b int
	var maxSt string
	for i := range s {
		//Odd length Palindrome
		a, b = i, i
		for a >= 0 && b < len(s) {
			if s[a] == s[b] {
				a--
				b++
			} else {
				break
			}
		}
		a, b = a+1, b-1
		if len(maxSt) < len(s[a:b+1]) {
			maxSt = s[a : b+1]
		}

		//Even length Palindrome
		a, b = i-1, i
		for a >= 0 && b < len(s) {
			if s[a] == s[b] {
				a--
				b++
			} else {
				break
			}
		}
		a, b = a+1, b-1
		if len(maxSt) < len(s[a:b+1]) {
			maxSt = s[a : b+1]
		}
	}
	return maxSt
}

// Runtime: 96ms; Runtime Percentile: 46.09%
// Memory: 2.7MB; Memory Percentile:  45.98%

// func longestPalindrome(s string) string {
// 	maxSt := ""
// 	maxlen := 0
// 	for i := range s {
// 		for j := i; j < len(s); j++ {
// 			if s[i] == s[j] {
// 				if maxlen < max(i, j) {
// 					if isPalindrome(i, j, &s) {
// 						maxlen = max(i, j)
// 						maxSt = s[i : j+1]
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return string(maxSt)
// }

// func isPalindrome(a int, b int, st *string) bool {
// 	for a <= b {
// 		if (*st)[a] != (*st)[b] {
// 			return false
// 		}
// 		a++
// 		b--
// 	}
// 	return true
// }
// func max(a int, b int) int {
// 	if a == b {
// 		return 1
// 	} else {
// 		return 1 + b - a
// 	}
// }

// Runtime: 788ms; Runtime Percentile: 11.05%
// Memory: 3.5MB; Memory Percentile:  37.95%

// func longestPalindrome(s string) string {
// 	maxSt := ""
// 	maxlen := 0
// 	for i := range s {
// 		fmt.Println("i:", i)
// 		for j := i; j < len(s); j++ {
// 			if s[i] == s[j] {
// 				if isPalindrome(i, j, &s) {
// 					if maxlen < max(i, j) {
// 						maxlen = max(i, j)
// 						maxSt = s[i : j+1]
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return string(maxSt)
// }

// func isPalindrome(a int, b int, st *string) bool {
// 	for a <= b {
// 		if (*st)[a] != (*st)[b] {
// 			return false
// 		}
// 		a++
// 		b--
// 	}
// 	return true
// }
// func max(a int, b int) int {
// 	if a == b {
// 		return 1
// 	} else {
// 		return 1 + b - a
// 	}
// }