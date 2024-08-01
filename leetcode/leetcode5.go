// 5. 最长回文子串
// 中等

// 提示
// 给你一个字符串 s，找到 s 中最长的
// 回文

// 子串
// 。

// 示例 1：

// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：

// 输入：s = "cbbd"
// 输出："bb"

// 提示：

// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
package main

// func main() {
// 	fmt.Println("== main begin ==")
// 	// s := "babad"
// 	// s := "aaabca"
// 	// s := "aacabdkacaa"
// 	// s := "a"
// 	s := "abccbd"
// 	ret := longestPalindrome(s)
// 	fmt.Println("== main end ==", ret)
// }

// // 回文需要对称
// func longestPalindrome(s string) string {
// 	var ret, tmpret string = "", ""
// 	var flag bool = false
// 	for i := 0; i < len(s); i++ {
// 		flag = false
// 		tmpret = ""
// 		for j := i + 1; j < len(s); j++ {
// 			if s[j] == s[i] {
// 				tmp := s[i : j+1]
// 				if len(tmp) > len(tmpret) {
// 					tmpret = tmp
// 				}
// 				if flag {
// 					break
// 				}
// 			} else {
// 				if !flag && len(tmpret) > 0 {
// 					break
// 				}
// 				flag = true
// 			}
// 		}
// 		if len(tmpret) > len(ret) {
// 			ret = tmpret
// 		}
// 	}
// 	if len(ret) == 0 {
// 		return s[0:1]
// 	}
// 	return ret
// }

// 回文需要对称
func longestPalindrome(s string) string {
	var ret string = ""
	var j, m, n int = 0, 0, 0
	for i := 0; i < len(s); i++ {
		j = 0
		for true {
			m = i - j
			n = i + j
			if m >= 0 && n < len(s) && s[m] == s[n] {
				tmp := s[m : n+1]
				if len(tmp) > len(ret) {
					ret = tmp
				}
			} else {
				break
			}
			j++
		}
		j = 0
		for true {
			m = i - j
			n = i + j + 1
			if m >= 0 && n < len(s) && s[m] == s[n] {
				tmp := s[m : n+1]
				if len(tmp) > len(ret) {
					ret = tmp
				}
			} else {
				break
			}
			j++
		}
	}
	if len(ret) == 0 {
		return s[0:1]
	}
	return ret
}
