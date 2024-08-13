package main

import (
	"fmt"
)

func test6(num1, num2 int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("=====err =", err)
		}
	}()
	fmt.Println("test6 div=", num1/num2, "i=", num2)
}

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		test6(100, i)
// 	}

// 	arr1 := make([]int, 4, 4)
// 	arr2 := arr1[2:4]
// 	arr2[0] = 1
// 	arr2[1] = 2
// 	fmt.Println("arr1", arr1, "len=", len(arr1), "cap=", cap(arr1)) // arr1 [0 0 1 2] len= 4 cap= 4
// 	fmt.Println("arr2", arr2, "len=", len(arr2), "cap=", cap(arr2)) // arr2 [1 2] len= 2 cap= 2
// 	arr2 = append(arr2, 3)
// 	arr2[0] = 3
// 	fmt.Println("arr1", arr1, "len=", len(arr1), "cap=", cap(arr1)) // arr1 [0 0 1 2] len= 4 cap= 4
// 	fmt.Println("arr2", arr2, "len=", len(arr2), "cap=", cap(arr2)) // arr2 [3 2 3] len= 3 cap= 4

// 	sort.Ints(arr1) //  使用sort包排序,默认为升序
// }
