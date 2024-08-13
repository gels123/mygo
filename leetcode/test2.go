package main

import "fmt"

type CarTool interface {
	Run()
}

type Car struct {
	Name string
}

func (*Car) Run() {
	fmt.Println("Car run")
}

type BiyadiCar struct {
	Car      // 匿名组合
	c    Car // 非匿名组合
	Name string
}

func (*BiyadiCar) Run() {
	fmt.Println("BiyadiCar run")
}

// func main() {
// 	// car := BiyadiCar{Name: "Biyadi", Car: Car{Name: "car"}}
// 	// fmt.Println(car.Name, car.Car.Name)
// 	// car.Run()     // 调用自身函数
// 	// car.c.Run()   // 调用非匿名组合的函数
// 	// car.Car.Run() // 调用匿名的父类函数

// 	car2 := Car{
// 		Name: "car2",
// 	}
// 	fmt.Println("car2=", car2.Name)
// 	car2.Run()
// }
