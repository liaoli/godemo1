package main

import "fmt"

func main() {
	demo2()
}

type Mover interface {
	move()
}

type dog struct{}

func (d dog) move() {
	fmt.Println("狗会动")
}

/*值接收者实现接口
值接收者和指针接收者实现接口的区别
从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。
*/
func demo1() {
	var x Mover
	var wangCai = dog{}
	x = wangCai

	var laifu = dog{}

	x = laifu

	x.move()
}

type Sayer interface {
	say()
}

type cat struct{}

func (d *cat) say() {
	fmt.Println("喵喵")
}

/*指针接收者实现接口
此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。
*/
func demo2() {

	var x Sayer

	var tom = cat{}

	// x = tom 报错

	x = &cat{} //x = &tom

	x = &tom
	x.say()
}
