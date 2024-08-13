// 1114. 按序打印
// 简单

// public class Foo {
//   public void first() { print("first"); }
//   public void second() { print("second"); }
//   public void third() { print("third"); }
// }
// 三个不同的线程 A、B、C 将会共用一个 Foo 实例。

// 线程 A 将会调用 first() 方法
// 线程 B 将会调用 second() 方法
// 线程 C 将会调用 third() 方法
// 请设计修改程序，以确保 second() 方法在 first() 方法之后被执行，third() 方法在 second() 方法之后被执行。

// 提示：

// 尽管输入中的数字似乎暗示了顺序，但是我们并不保证线程在操作系统中的调度顺序。
// 你看到的输入格式主要是为了确保测试的全面性。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出："firstsecondthird"
// 解释：
// 有三个线程会被异步启动。输入 [1,2,3] 表示线程 A 将会调用 first() 方法，线程 B 将会调用 second() 方法，线程 C 将会调用 third() 方法。正确的输出是 "firstsecondthird"。
// 示例 2：

// 输入：nums = [1,3,2]
// 输出："firstsecondthird"
// 解释：
// 输入 [1,3,2] 表示线程 A 将会调用 first() 方法，线程 B 将会调用 third() 方法，线程 C 将会调用 second() 方法。正确的输出是 "firstsecondthird"。

// 提示：
// nums 是 [1, 2, 3] 的一组排列

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Foo struct {
	flag int32
	chs  [3]chan int
	wg   sync.WaitGroup
}

func (s *Foo) init() {
	if atomic.LoadInt32(&s.flag) == 0 {
		atomic.StoreInt32(&s.flag, 1)
		s.chs[0] <- 1
	}
}

func (s *Foo) First() {
	s.init()
	<-s.chs[0]
	fmt.Print("first")
	s.wg.Done()
	s.chs[1] <- 1
}

func (s *Foo) Second() {
	s.init()
	<-s.chs[1]
	fmt.Print("second")
	s.wg.Done()
	s.chs[2] <- 1
}

func (s *Foo) Third() {
	s.init()
	<-s.chs[2]
	fmt.Print("third")
	s.wg.Done()
}

// func main() {
// 	fmt.Println("==main begin==")
// 	foo := &Foo{chs: [3]chan int{make(chan int), make(chan int), make(chan int)}}
// 	var arr [3]interface{}
// 	arr[0] = foo.First
// 	arr[1] = foo.Second
// 	arr[2] = foo.Third

// 	sequece := []int{1, 2, 3}
// 	foo.wg.Add(3)
// 	for _, v := range sequece {
// 		v--
// 		if arr[v] != nil {
// 			f := reflect.ValueOf(arr[v])
// 			if f.IsValid() && f.Kind() == reflect.Func {
// 				// fmt.Println("call v=", v, f.Kind(), f.Kind() == reflect.Func)
// 				params := []reflect.Value{}
// 				go f.Call(params)
// 			}
// 		}
// 	}
// 	foo.wg.Wait()
// 	fmt.Println("")
// 	fmt.Println("==main end==")
// }
