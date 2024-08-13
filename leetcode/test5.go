package main

import "fmt"

func test5(nums ...int) {
	fmt.Println("test5==", len(nums), nums, nums[4])
}

// func main() {
// 	var aa interface{}
// 	aa = 100
// 	fmt.Println("==========", aa.(int))
// 	nums := []int{1, 2, 3, 4, 5}
// 	test5(nums...) // ...散开参数, append(slices, nums...)
// }
