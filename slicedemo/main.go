package main

import "fmt"

func main() {
	sliceDemo13()
}

func sliceDemo1() {
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化

	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
}

func sliceDemo2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}

/*对切片在执行切片表达式
 */
func sliceDemo3() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]//简单切片表达式，通过一个数组得到一个切片
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

/*
*为了方便起见，可以省略切片表达式中的任何索引。
*省略了low则默认为0；
*省略了high则默认为切片操作数的长度
 */
func sliceDemo4() {
	a := [5]int{1, 2, 3, 4, 5}
	s1 := a[2:] // 等同于 a[2:len(a)]
	s2 := a[:3] // 等同于 a[0:3]
	s3 := a[:]  // 等同于 a[0:len(a)]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

/*完整切片表达式
 */
func sliceDemo5() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:3] //完整切片表达式  a[low : high : max] high <= max <= len(a)
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}

/*
使用make()函数构造切片
我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：

make([]T, size, cap)
其中：

T:切片的元素类型
size:切片中元素的数量
cap:切片的容量
*/
func sliceDemo6() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}

/*判断切片是否为空
 *切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。
 *切片唯一合法的比较操作是和nil比较。
 *一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
 *但是我们不能说一个长度和容量都是0的切片一定是nil
 */
func sliceDemo7() {

	var s1 []int

	s2 := []int{}

	s3 := make([]int, 0)

	fmt.Printf("s1:%v len(s1):%v cap(s1):%v s1 == nil %v\n", s1, len(s1), cap(s1), s1 == nil)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v s2 == nil %v\n", s2, len(s2), cap(s2), s2 == nil)
	fmt.Printf("s3:%v len(s3):%v cap(s3):%v s3 == nil %v\n", s3, len(s3), cap(s3), s3 == nil)
	// s1:[] len(s1):0 cap(s1):0 s1 == nil true
	// s2:[] len(s2):0 cap(s2):0 s2 == nil false
	// s3:[] len(s3):0 cap(s3):0 s3 == nil false
}

/*切片的赋值拷贝
 */
func sliceDemo8() {

	s1 := make([]int, 3)
	s2 := s1 //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}

/*切片遍历 与数组一样
 */
func sliceDemo9() {

	s := []int{1, 32, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for index, value := range s {
		fmt.Printf("index:%v --> value:%v\n", index, value)
	}
}

/*append()方法为切片添加元素
 *Go语言的内建函数append()可以为切片动态添加元素。
 *可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）
 */
func sliceDemo10() {

	var s []int
	//注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化
	s = append(s, 1)
	fmt.Printf("s:%v len(s):%v cap(s):%v \n", s, len(s), cap(s)) //s:[1] len(s):1 cap(s):1
	s = append(s, 2, 3, 4, 5)
	fmt.Printf("s:%v len(s):%v cap(s):%v \n", s, len(s), cap(s)) //s:[1 2 3 4 5] len(s):5 cap(s):6
	s2 := []int{6, 7, 8}

	s = append(s, s2...) //可以添加另一个切片中的元素（后面加…）

	fmt.Printf("s:%v len(s):%v cap(s):%v \n", s, len(s), cap(s)) //s:[1 2 3 4 5 6 7 8] len(s):8 cap(s):12

}

/*每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，
此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，
所以我们通常都需要用原变量接收append函数的返回值。
*/
func sliceDemo11() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

/*使用copy()函数复制切片
由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。

Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下
copy(destSlice, srcSlice []T)
其中：

srcSlice: 数据来源切片
destSlice: 目标切片
*/
func sliceDemo12() {

	srcS := []int{1, 2, 4, 5, 6}

	destS := make([]int, 5, 5)

	copy(destS, srcS)

	destS[3] = 77

	fmt.Println(srcS)

	fmt.Println(destS)

}

/*从切片中删除元素
Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。
要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
*/
func sliceDemo13() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	s = append(s[:2], s[3:]...)
	fmt.Println(s)

}
