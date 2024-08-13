package main

import (
	"fmt"
)

type Test4 struct {
	name string
	age  int
}

func (Test4) F1(num1, num2 int) {
	fmt.Println("==Test4 f1()==", num1, num2)
}

// func main() {
// 	t0 := Test4{
// 		name: "xiaoming",
// 		age:  14,
// 	}
// 	var t interface{} = t0
// 	t_value := reflect.ValueOf(t) // 获取值
// 	// reflect.TypeOf=类型,指类名   Kind=指数据类型struct
// 	fmt.Println("ValueOf t=", t_value, "TypeOf t=", reflect.TypeOf(t), t_value.Kind()) // ValueOf t= {xiaoming 14} TypeOf t= main.Test4 struct
// 	t_f1 := t_value.MethodByName("F1")                                                 //  大写开头的方法才能获取到
// 	fmt.Println("f1 kind=", t_f1.Kind())
// 	if t_f1.Kind() == reflect.Func && t_f1.IsValid() {
// 		t_f1.Call([]reflect.Value{reflect.ValueOf(11), reflect.ValueOf(12)})
// 	}
// }
