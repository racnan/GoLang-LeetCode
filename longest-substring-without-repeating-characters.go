// Runtime: 0ms; Runtime Percentile: 100.00%
// Memory: 2.6MB; Memory Percentile:  99.94%

func lengthOfLongestSubstring(s string) int {
	table := [128]int{}
	i, maxlen := 0, 0
	for j, ch := range s {
		for table[ch] == 1 {
			table[s[i]] = 0
			i++
			if maxlen < max(i, j) {
				maxlen = max(i, j)
			}
		}
		table[ch] += 1
		if maxlen < max(i, j) {
			maxlen = max(i, j)
		}
	}
	return maxlen
}

func max(a int, b int) int {
	if a == b {
		return 1
	} else {
		return 1 + b - a
	}
}

// Runtime: 12ms; Runtime Percentile: 41.15%
// Memory: 2.6MB; Memory Percentile:  99.94%

// func lengthOfLongestSubstring(s string) int {
// 	table := [128]int{}
// 	i, maxlen := 0, 0
// 	for _, ch := range s {
// 		for table[ch] == 1 {
// 			table[s[i]] = 0
// 			i++
// 			if maxlen < max(&table) {
// 				maxlen = max(&table)
// 			}
// 		}
// 		table[ch] += 1
// 		if maxlen < max(&table) {
// 			maxlen = max(&table)
// 		}
// 	}
// 	return maxlen
// }

// func max(tb *[128]int) int {
// 	count := 0
// 	for _, ones := range *tb {
// 		if ones == 1 {
// 			count++
// 		}
// 	}
// 	return count
// }

// Runtime: 1292ms; Runtime Percentile: 5.14%
// Memory: 6.5MB; Memory Percentile: 19.13%

// func lengthOfLongestSubstring(s string) int {
// 	str := ""
// 	max := 0
// 	match := false
// 	for i := 0; i < len(s); i++ {
// 		str += string(s[i])
// 		for j := i + 1; j < len(s); j++ {
// 			for _, ch := range str {
// 				if string(s[j]) == string(ch) {
// 					match = true
// 					break
// 				}
// 			}
// 			if match {
// 				break
// 			}
// 			str += string(s[j])
// 		}
// 		match = false
// 		if len(str) > max {
// 			max = len(str)
// 		}

// 		str = ""
// 	}
// 	return max
// }
