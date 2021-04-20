package main

import "fmt"

func main() {

	w := haier{
		dryer{},
	}

	w.wash()
	w.dry()
	w.dryer.dry()

	var wm WashingMachine

	wm = w

	wm.wash()
	wm.dry()
}

/*
并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

*/

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct{}

type haier struct {
	dryer
}

func (d dryer) dry() {
	fmt.Println("甩干")
}

func (h haier) wash() {
	fmt.Println("洗一洗")
}
