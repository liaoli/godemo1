package main

func mian() {

}

/*
并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

*/

type WashingMachine interface {
	wash()
	dry()
}
