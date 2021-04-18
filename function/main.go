package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
Go语言中定义函数使用func关键字，具体格式如下：

func 函数名(参数)(返回值){
    函数体
}
其中：

函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
函数体：实现指定功能的代码块。
*/
func main() {
	panicRecoverDemo()
}

/*定义函数类型
*以下语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
*简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型
 */
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

/*函数类型与变量
 */
func funcDemo1() {
	var c calculation

	c = add

	result := c(1, 2)
	fmt.Printf("type of c is %T\n", c)
	fmt.Println(result)

	f := sub

	fmt.Printf("type of f is %T\n", f)
	fmt.Println(f(1, 2))
}

/*高阶函数
高阶函数分为函数作为参数和函数作为返回值两部分。
*/

/*函数作为参数
 */
func cal(x, y int, op func(int, int) int) int {
	return op(x, y)
}

/*函数作为返回值
 */
func do(s string) (func(int, int) int, error) {

	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func gaoJieFuncDemo() {
	f1 := cal

	fmt.Printf("type of f1 is %T\n", f1)
	fmt.Println(f1(1, 2, add))
	fmt.Println(f1(1, 2, sub))

	f2 := do

	f3, _ := f2("+")

	fmt.Printf("type of f3 is %T\n", f3)
	fmt.Println(f3(1, 2))
}

/*匿名函数和闭包
函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。
匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

func(参数)(返回值){
    函数体
}
匿名函数因为没有函数名，所以没办法像普通函数那样调用，
所以匿名函数需要保存到某个变量或者作为立即执行函数:
*/
func noNameFuc() {

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2)) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	result := func(x, y int) int {
		return x - y
	}(1, 2)

	fmt.Println(result)

}

/*闭包
闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
*/

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixfunc(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

func calc(base int) (func(int) int, func(int) int) {

	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func closureDemo() {
	/*
	   变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。
	    在f的生命周期内，变量x也一直有效。 闭包进阶示例1：
	*/
	var f = adder()

	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(30))

	f1 := adder()

	fmt.Println(f1(100))
	fmt.Println(f1(200))

	jpgFunc := makeSuffixfunc(".jpg")
	exeFunc := makeSuffixfunc(".exe")

	fmt.Println(jpgFunc("test"))
	fmt.Println(exeFunc("test"))

	f2, f3 := calc(10)

	fmt.Println(f2(1), f3(1)) //11 10

	fmt.Println(f2(3), f3(2)) //13 11

	fmt.Println(f2(5), f3(7)) //16 9

}

/*defer语句
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，
将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等
defer执行时机
在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前
*/
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func deferDemo1() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func deferDemo2() { //defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	x := 1
	y := 2
	defer calc1("AA", x, calc1("A", x, y))
	x = 10
	defer calc1("BB", x, calc1("B", x, y)) // A 1 2 3  B 10 2 12 BB 10 12 22 AA 1 3 4
	y = 20
}

func calc1(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/*Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。
panic可以在任何地方引发，但recover只有在defer调用的函数中有效。
*/

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcB1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic in B1")
}

func funcC() {
	fmt.Println("func C")
}

func panicRecoverDemo() {
	funcA()
	/*
		func A
		panic: panic in B

		goroutine 1 [running]:
		main.funcB(...)
		        /Users/mac/godemo/godemo1/function/main.go:266
		main.panicRecoverDemo()
		        /Users/mac/godemo/godemo1/function/main.go:275 +0x96
		main.main()
		        /Users/mac/godemo/godemo1/function/main.go:23 +0x25
		exit status 2
	*/
	// funcB()
	funcB1()
	funcC()

}
