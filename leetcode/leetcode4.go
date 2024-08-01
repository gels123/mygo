// 4. 寻找两个正序数组的中位数
// 困难

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

// 算法的时间复杂度应该为 O(log (m+n)) 。

// 示例 1：

// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：

// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

// 提示：

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

package main

// func main() {
// 	fmt.Println("== main begin ==")
// 	nums1 := []int{1, 2}
// 	nums2 := []int{3, 4}
// 	ret := findMedianSortedArrays(nums1, nums2)
// 	fmt.Println("== main end ==", ret)
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1)-1, len(nums2)-1
	cnt := 0
	total := len1 + len2 + 2
	var p1, p2 int = -1, -1
	var p3, p4 int = len1 + 1, len2 + 1
	var m, n int = 0, 0
	for cnt < total {
		if p1 < len1 && p2 < len2 {
			if nums1[p1+1] < nums2[p2+1] {
				p1++
				m = nums1[p1]
			} else {
				p2++
				m = nums2[p2]
			}
			cnt++
		} else if p1 < len1 {
			p1++
			m = nums1[p1]
			cnt++
		} else if p2 < len2 {
			p2++
			m = nums2[p2]
			cnt++
		}
		if p3 > 0 && p4 > 0 {
			if nums1[p3-1] > nums2[p4-1] {
				p3--
				n = nums1[p3]
			} else {
				p4--
				n = nums2[p4]
			}
			cnt++
		} else if p3 > 0 {
			p3--
			n = nums1[p3]
			cnt++
		} else if p4 > 0 {
			p4--
			n = nums2[p4]
			cnt++
		}
	}
	return ((float64)(m + n)) / 2
}
