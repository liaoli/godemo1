package main

import "fmt"

func main() {

	var s Sayer
	var m Mover
	var a animal

	c := cat{
		"Tom",
	}

	s = c
	m = c
	a = c

	s.say()
	m.move()
	a.say()
	a.move()

}

/*接口嵌套
接口与接口间可以通过嵌套创造出新的接口。


*/
// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
/*
嵌套得到的接口的使用与普通接口一样，
这里我们让cat实现animal接口：
*/
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("跳桌子上")
}
