package main

import "fmt"

type MyF func(a, b int)

func test8_f1(a, b int) {
	fmt.Println("test8_f1", a+b)
}

func test8_f2(a, b int) {
	fmt.Println("test8_f2", a-b)
}

// func main() {

// 	fmt.Println("===main begin==")
// 	arr := []MyF{test8_f1, test8_f2}
// 	for _, f := range arr {
// 		f(1, 2)
// 	}
// 	fmt.Println("===main end==")
// }
