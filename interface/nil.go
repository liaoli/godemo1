package main

import "fmt"

/*
空接口的定义
空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

空接口类型的变量可以存储任意类型的变量。
*/
func main() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)

	show(100)

	show("haha")

	var u struct {
		int
		string
	}

	show(u)

	demo()

	duanyan(x)
	duanyan1()
}

/*
空接口的应用
空接口作为函数的参数
使用空接口实现可以接收任意类型的函数参数。
*/
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

/*
空接口作为map的值
使用空接口实现可以保存任意值的字典。
*/
func demo() {

	var stu = make(map[string]interface{})

	stu["name"] = "xxx"
	stu["age"] = 10
	stu["married"] = false

	show(stu)

}

/*类型断言
 */

func duanyan1() {

	var a interface{}

	a = false

	v, ok := a.(bool)

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}

func duanyan(a interface{}) {

	switch v := a.(type) {
	case string:
		fmt.Printf("a is a sting ,value is %s", v)
	case int:
		fmt.Printf("a is a int ,value is %d", v)
	case bool:
		fmt.Printf("a is a bool ,value is %v", v)
	default:
		fmt.Println("unsupport type!")
	}
}
