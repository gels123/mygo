package main

// func main() {
// 	ch := make(chan int, 3)
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	close(ch) // 无此行将 fatal error: all goroutines are asleep - deadlock!
// 	ch = nil
// 	select {
// 	case ch <- 4:
// 	default:
// 		{
// 			if ch == nil {
// 				fmt.Println("default1")
// 			} else {
// 				fmt.Println("default2")
// 			}
// 		}
// 	}
// 	// ch <- 4   // 有此行将 panic: send on closed channel
// 	// for num := range ch {
// 	// 	fmt.Println(num)
// 	// }
// }
