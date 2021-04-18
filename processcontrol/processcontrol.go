package main

import "fmt"

func main() {
	switchDemo4()
}

func switchDemo1() {
	finger := 3

	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效输入")
	}

}

func switchDemo2() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
}

func switchDemo3() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

//fallthrough语法可以执行满足条件的case的下一个case
func switchDemo4() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func forDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

//无限循环
func forDemo4() {
	for {
		fmt.Println("*")

	}
}

func ifDemo() {

	//此score 的作用域只在if 语句块起作用
	if score := 65; score > 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else if score > 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

}
