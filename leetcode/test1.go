package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("===main begin==")
	var s string = "abcdefg"

	var wg sync.WaitGroup
	wg.Add(2)
	c1, c2 := make(chan int, 0), make(chan int, 0)
	go func() {
		for {
			i := <-c1
			fmt.Println("co1=", i, string(s[i]))
			i++
			if i >= len(s)-1 {
				break
			}
			c2 <- i
		}
		wg.Done()
	}()

	go func() {
		for {
			i := <-c2
			fmt.Println("co2=", i, string(s[i]))
			i++
			c1 <- i
			if i >= len(s)-2 {
				break
			}
		}
		wg.Done()
	}()
	c1 <- 0
	wg.Wait()

	fmt.Print("===main end===")
}
