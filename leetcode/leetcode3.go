// 3. 无重复字符的最长子串
// 中等
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长
// 子串
//  的长度。

// 示例 1:

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2p

// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 提示：

// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成

package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("== main begin==")
// 	var str string = "abcabcbb"
// 	ret := lengthOfLongestSubstring(str)
// 	fmt.Println("== main end==", ret)
// }

// 递归 or 动态规划dp or 滑动窗口
func lengthOfLongestSubstring(s string) int {
	fmt.Println("s=", s, "len=", len(s), "sub=")
	ret := 0
	dp := make(map[int][]string, 0)
	for i := 0; i < len(s); i++ {
		arr, ok := dp[i-1]
		if ok {
			newdp := make([]string, 0)
			c := s[i : i+1]
			newdp = append(newdp, c)
			if len(c) > ret {
				ret = len(c)
			}
			for _, v := range arr {
				c = s[i : i+1]
				var find bool = false
				for ii := 0; ii < len(v); ii++ {
					if v[ii] == c[0] {
						find = true
						break
					}
				}
				if !find {
					str2 := v + c
					newdp = append(newdp, str2)
					if len(str2) > ret {
						ret = len(str2)
					}
				}
			}
			dp[i] = newdp
		} else {
			dp[i] = make([]string, 0)
			c := s[i : i+1]
			dp[i] = append(dp[i], c)
			if len(c) > ret {
				ret = len(c)
			}
		}
		delete(dp, i-2)
	}
	return ret
}
