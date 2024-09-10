// 45. 跳跃游戏 II
// 中等

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// 示例 1:

// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 示例 2:

// 输入: nums = [2,3,0,1,4]
// 输出: 2

// 提示:

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// 题目保证可以到达 nums[n-1]

package main

import (
	"fmt"
)

// 贪心算法
func jump2(nums []int) int {
	var ret int = 0
	l := len(nums) - 1
	if l == 0 {
		return ret
	}
	if nums[0] >= l {
		ret++
		return ret
	}
	max, maxi, tmp := 0, 0, 0
	for i := 0; i <= l; {
		max, maxi = 0, 0
		for j := i + 1; j <= min(l, i+nums[i]); j++ {
			tmp = j + nums[j]
			if maxi == 0 || tmp >= max {
				max, maxi = tmp, j
			}
		}
		if maxi == 0 {
			break
		} else {
			ret++
			i = maxi
			if i >= l {
				break
			}
			if max >= l {
				ret++
				break
			}
		}
	}
	return ret
}

// 类似动态规划
func jump(nums []int) int {
	l := len(nums) - 1
	if l == 0 {
		return 0
	}
	if nums[0] >= l {
		return 1
	}
	step := make([]int, l+1)
	for i := 0; i <= l; i++ {
		step[i] = 0
	}
	// fmt.Println("step len=", len(step), "camp=", cap(step), "step=", step)
	var i, j, maxj int
	for i = 0; i <= l; i++ {
		maxj = i + nums[i]
		if maxj > l {
			maxj = l
		}
		for j = i + 1; j <= maxj; j++ {
			if step[j] == 0 {
				step[j] = step[i] + 1
				if j == l {
					return step[l]
				}
			}
		}
	}
	// fmt.Println("end step=", step)
	return step[l]
}

func main() {
	fmt.Println("== main begin ==")
	// var nums []int = []int{2, 3, 1, 0}
	var nums []int = []int{2, 3, 0, 1, 4}
	ret := jump(nums)
	fmt.Println("jump ret=", ret)

	fmt.Println("== main end ==")
}
