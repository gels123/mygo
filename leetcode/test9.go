// 贪心算法测试
// 假设1元、2元、5元、10元、20元、50元、100元的纸币分别有c0, c1, c2, c3, c4, c5, c6张。现在要用这些钱来支付K元，至少要用多少张纸币？
package main

import "fmt"

type Money struct {
	val, cnt int
}

func main() {
	fmt.Println("== main begin ==")
	money := []Money{{1, 20}, {2, 20}, {5, 20}, {10, 20}, {50, 20}, {100, 20}}
	pay := 532
	amount := 0
	for _, v := range money {
		amount += v.val * v.cnt
	}
	if pay > amount {
		fmt.Println("pay > amount", pay, amount)
		return
	}
	tmppay := pay
	for tmppay > 0 {

	}

	fmt.Println("== main end ==")
}

func test_merge() {

}
