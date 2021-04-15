package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// fmt.Println("Hello World")

	// fmt.Print("a", "b", "c")

	// fmt.Println()

	// fmt.Println("a", "b", "c")

	// fmt.Printf("hello world")

	//ConstantDeclarationAndInitialization()
	dataType()
}

func dataType() {

	// var num int = 10
	// fmt.Printf("num=%v 类型：%T \n", num, num)
	// var num int8 = 120
	// fmt.Printf("num=%v size of 类型%T  is %v \n", num, num, unsafe.Sizeof(num))

	// var num int8 = 129 //constant 129 overflows int8
	// fmt.Printf("num=%v size of 类型%T  is %v \n", num, num, unsafe.Sizeof(num))

	// var num int16 = 129 //constant 129 overflows int8
	// fmt.Printf("num=%v size of 类型%T  is %v \n", num, num, unsafe.Sizeof(num))

	// var num int32 = 129 //constant 129 overflows int8
	// fmt.Printf("num=%v size of 类型%T  is %v \n", num, num, unsafe.Sizeof(num))

	var num int64 = 129 //constant 129 overflows int8
	fmt.Printf("num=%v size of 类型%T  is %v \n", num, num, unsafe.Sizeof(num))

	var n = 33

	fmt.Printf("num=%v size of %T is %v \n", n, n, unsafe.Sizeof(n))

	n1 := 22

	fmt.Printf("n1 = %v \n", n1) // %v 原样输出

	fmt.Printf("n1 十进制= %d \n", n1) // %d 十进制输出

	fmt.Printf("n1  二进制= %b \n", n1) // %b 二进制输出

	fmt.Printf("n1 八进制= %o \n", n1) // %o 八进制输出

	fmt.Printf("n1 十六进制= %x \n", n1) // %x 十六机制输出

}

func ConstantDeclarationAndInitialization() {
	const pi = 3.1415926
	fmt.Println(pi)

	//一次定义多个常量 只给第一个赋值，后面的常量与第一个常量的值一样
	const (
		first = 1
		socend
		third
		forth
	)

	fmt.Println(first, socend, third, forth)

	//关键字 iota 遇到 const 之后值重置为0，

	const cc = iota

	fmt.Println("cc = ", cc)

	const (
		id1 = iota
		id2
		id3
	)

	fmt.Println(id1, id2, id3)
	// fmt.Println(iota) 报错 必须与 const 初始化的时候才能用
	const aa = iota

	fmt.Println("aa = ", aa)

	//使用_跳过某些值
	const (
		n1 = iota // 0
		n2        //1
		_
		n4 //3
	)
	fmt.Println(n1, n2, n4)
	//iota声明中间插队
	const (
		m1 = iota // 0
		m2 = 100
		m3 = iota //2
		m4        // 3
	)
	fmt.Println(m1, m2, m3, m4)
	//多个iota定义在一行
	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
	fmt.Println(a, b, c, d, e, f)
}

/*
变量声明和初始化
*/
func VariableDeclarationAndInitialization() {
	var a int = 1

	var b = 2

	fmt.Println(a, b)

	fmt.Println("a = ", a, "b = ", b)

	fmt.Printf("a = %d,b = %d", a, b)

	fmt.Printf("a = %d,b = %d ,a = %T, b = %T\n", a, b, a, b)

	i, j, k := 1, 2, "fff"

	fmt.Println(i, j, k)

	var (
		f = 1
		g = 2
	)
	fmt.Println(f, g)
	var aa, bb string

	aa = "33"
	bb = "cc"

	fmt.Println(aa, bb)

	var integer int

	fmt.Println(integer)

	var s string

	fmt.Println(s)

	fmt.Println(testUnderLine())

	var name, _ = testUnderLine()

	fmt.Println(name)
}

func testUnderLine() (string, int) {
	return "mumu", 33
}
