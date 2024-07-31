// 1. 两数之和
// 简单
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

// 你可以按任意顺序返回答案。

// 示例 1：

// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
// 示例 2：

// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
// 示例 3：

// 输入：nums = [3,3], target = 6
// 输出：[0,1]

// 提示：

// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// 只会存在一个有效答案

// 进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？

package main

// func main() {
// 	fmt.Println("==main begin==")
// 	// nums := []int{3, 2, 4}
// 	nums := []int{3, 3}
// 	target := 6
// 	ret := twoSum(nums[:], target)
// 	fmt.Println("==main end ret=", ret)
// }

// func twoSum(nums []int, target int) []int {
// 	var ret = []int{0, 0}
// 	for i := 0; i < len(nums); i++ {
// 		fmt.Println("i=", i, "v=", nums[i])
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				ret[0] = i
// 				ret[1] = j
// 				return ret
// 			}
// 		}
// 	}
// 	return ret
// }

func twoSum(nums []int, target int) []int {
	var ret = []int{0, 0}
	var m = make(map[int][]int, 0)
	for i := 0; i < len(nums); i++ {
		// fmt.Println("i=", i, "v=", nums[i])
		arr, ok := m[nums[i]]
		if ok {
			m[nums[i]] = append(arr, i)
		} else {
			m[nums[i]] = []int{i}
		}
	}
	for v, k := range m {
		// fmt.Println("k=", k, "v=", v)
		n := target - v
		karr, ok := m[n]
		if ok {
			for _, k1 := range k {
				for _, k2 := range karr {
					if k1 != k2 {
						ret[0] = k1
						ret[1] = k2
						return ret
					}
				}
			}
		}
	}
	return ret
}
