package main

import "fmt"

func main() {
	pointDemo4()
}

/*指针地址和指针类型
 */
func pointDemo1() {
	a := 10

	b := &a

	fmt.Printf("a= %d ptr : %p \n", a, &a)
	fmt.Printf("b= %v type : %T ptr:%p\n", b, b, &b)
}

/*指针取值
 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
对变量进行取地址（&）操作，可以获得这个变量的指针变量。
指针变量的值是指针地址。
对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
*/
func pointDemo2() {
	a := 10

	b := &a

	c := *b

	fmt.Printf("a= %d ptr : %p \n", a, &a)
	fmt.Printf("b= %v type : %T ptr:%p\n", b, b, &b)

	fmt.Printf("c= %v type : %T ptr:%p\n", c, c, &c)

	// a= 10 ptr : 0xc000016088
	// b= 0xc000016088 type : *int ptr:0xc00000e028
	// c= 10 type : int ptr:0xc000016090
}

/*指针传值示例

 */
func pointDemo3() {

	a := 10
	modify1(a)
	fmt.Println(a) //10
	modify2(&a)
	fmt.Println(a) //100
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

/*new和make
对于引用类型的变量，我们在使用的时候不仅要声明它，
还要为它分配内存空间，否则我们的值就没办法存储。
new是一个内置的函数，它的函数签名如下：

func new(Type) *Type
其中，

Type表示类型，new函数只接受一个参数，这个参数是一个类型
*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
*/
func pointDemo4() {

	// var a *int
	//会报错引发panic 对于引用类型的变量，我们在使用的时候不仅要声明它，
	//还要为它分配内存空间，否则我们的值就没办法存储。
	// *a = 100

	a := new(int)
	fmt.Printf("%T ,%d\n", a, *a) //*int ,0

	// var b map[string]int
	// b["沙河娜扎"] = 100//会报错引发panic b 没有初始化
	// fmt.Println(b)

	// 	make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
	// 而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，
	// 所以就没有必要返回他们的指针了。make函数的函数签名如下：
	// func make(t Type, size ...IntegerType) Type
	// make函数是无可替代的，我们在使用slice、map以及channel的时候，
	// 都需要使用make进行初始化，然后才可以对它们进行操作。

	var b map[string]int

	b = make(map[string]int, 10)
	b["广东省"] = 1
	fmt.Println(b)

}

// new与make的区别
// 二者都是用来做内存分配的。
// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
// 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
