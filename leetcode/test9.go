// 贪心算法测试
// 假设1元、2元、5元、10元、20元、50元、100元的纸币分别有c0, c1, c2, c3, c4, c5, c6张。现在要用这些钱来支付K元，至少要用多少张纸币？
package main

import (
	"fmt"
)

type Money struct {
	val, cnt int
}

func main() {
	fmt.Println("== main begin ==")
	money := []Money{{1, 20}, {2, 20}, {5, 20}, {10, 20}, {50, 20}, {100, 20}}
	var result []Money
	pay := 532
	for i := len(money) - 1; i >= 0; i-- {
		cel := money[i]
		cnt_ := min(cel.cnt, pay/cel.val)
		if cnt_ > 0 {
			pay -= cnt_ * cel.val
			result = append(result, Money{cel.val, cnt_})
		}
	}
	if pay == 0 {
		fmt.Println("find ok result=", result)
	} else {
		fmt.Println("find fail result=", result)
	}
	fmt.Println("== main end ==")
}

func test_merge() {

}

func test_merge2() {

}

func test_main1() {

}
