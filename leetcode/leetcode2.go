// 2. 两数相加
// 中等

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例 1：

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：

// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：

// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]

// 提示：

// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零

package main

// func main() {
// 	fmt.Println("==main begin ==")

// 	l1 := &ListNode{
// 		Val: 2,
// 		Next: &ListNode{
// 			Val: 4,
// 			Next: &ListNode{
// 				Val:  3,
// 				Next: nil,
// 			},
// 		},
// 	}
// 	l2 := &ListNode{
// 		Val: 5,
// 		Next: &ListNode{
// 			Val: 6,
// 			Next: &ListNode{
// 				Val:  4,
// 				Next: nil,
// 			},
// 		},
// 	}
// 	ret := addTwoNumbers(l1, l2)
// 	fmt.Println("==main end ==", ret)
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode = nil
	var ret *ListNode = nil
	var tmp = 0
	for true {
		if l1 == nil && l2 == nil {
			break
		}
		if ret == nil {
			ret = &ListNode{Val: 0, Next: nil}
			root = ret
		} else {
			ret.Next = &ListNode{Val: 0, Next: nil}
			ret = ret.Next
		}
		if l1 != nil && l2 != nil {
			tmp += l1.Val + l2.Val
			ret.Val = tmp % 10
			tmp = tmp / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			tmp += l1.Val
			ret.Val = tmp % 10
			tmp = tmp / 10
			l1 = l1.Next
		} else {
			tmp += l2.Val
			ret.Val = tmp % 10
			tmp = tmp / 10
			l2 = l2.Next
		}
	}
	if tmp > 0 {
		ret.Next = &ListNode{Val: tmp, Next: nil}
	}
	return root
}
