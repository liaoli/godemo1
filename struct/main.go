package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

/*
问题

1.面试题
m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

 2.	json.Unmarshal([]byte(str), c1)
*/
func main() {
	imitateExtends()
}

/*结构体的定义
使用type和struct关键字来定义结构体，具体代码格式如下：

type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
其中：

类型名：标识自定义结构体的名称，在同一个包内不能重复。
字段名：表示结构体字段名。结构体中的字段名必须唯一。
字段类型：表示结构体字段的具体类型。
*/
type person struct {
	name string
	age  int
	city string
}

/*结构体初始化 赋值
var 结构体实例 结构体类型
*/
func structDemo1() {

	var p1 person
	fmt.Println(p1)
	p1.name = "丽丽"
	p1.age = 22
	p1.city = "深圳"

	fmt.Printf("p1 = %v\n", p1)

	fmt.Printf("p1 = %#v\n", p1)
}

/*面试题 说出打印结果
 */

func test() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

/*匿名结构体
 */
func structDemo2() {

	var user struct {
		name string
		age  int
	}

	fmt.Printf("user = %v\n", user)

	user.name = "hh"
	user.age = 13
	fmt.Printf("user = %#v\n", user)

}

/*指针类型结构体*/
func structDemo3() {
	var p2 = new(person)
	fmt.Printf("p2 = %T\n", p2)
	fmt.Printf("p2 = %#v\n", p2) //&main.person{name:"", age:0, city:""}
	p2.name = "里奥哈哈"
	p2.age = 33
	p2.city = "深圳"
	fmt.Printf("p2 = %v\n", p2) //p2 = &{里奥哈哈 33 深圳}

	fmt.Printf("p2 = %v\n", *p2) //p2 = {里奥哈哈 33 深圳}
}

/*取结构体的地址实例化
使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
*/
func structDemo4() {

	p := &person{}
	fmt.Printf("p = %T\n", p)
	fmt.Printf("p = %#v\n", p)

	p.name = "杨小胖"
	p.age = 33
	p.city = "广州"
	fmt.Printf("p = %#v\n", p)

	fmt.Printf("p = %#v\n", *p)
}

/*结构体初始化
使用键值对初始化
*/
func structDemo5() {
	// 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	p := person{
		name: "蒙查查",
		city: "东莞",
		age:  22,
	}

	fmt.Printf("p = %#v\n", p)

	// 也可以对结构体指针进行键值对初始化，例如：
	p1 := &person{
		name: "李大大",
		city: "广州",
		age:  22,
	}

	fmt.Printf("p = %#v\n", p1)

	fmt.Printf("p = %#v\n", *p1)

	// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

	p2 := person{
		name: "嘻嘻嘻",
	}

	fmt.Printf("p = %#v\n", p2)

	p3 := &person{
		name: "嘻嘻嘻",
	}

	fmt.Printf("p = %#v\n", p3)
}

/*使用值的列表初始化
初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
*/
func structDemo6() {

	p := &person{
		"杨哈哈",
		33,
		"永州",
	}

	fmt.Printf("p = %#v\n", p)
	// 	使用这种格式初始化时，需要注意：
	// 必须初始化结构体的所有字段。
	// 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 该方式不能和键值初始化方式混用。
}

/*结构体内存布局
 */
type memory struct {
	a int8
	b int8
	c int8
	d int8
}

func memoryDemo() {

	n := memory{
		1, 2, 3, 4,
	}

	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

}

/*构造函数
Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
*/
func newPerson(name, city string, age int) *person {
	return &person{
		name,
		age,
		city,
	}
}

//调用构造函数

func methodDemo() {

	p := newPerson("ccc", "永州", 33)

	fmt.Printf("p = %#v\n", *p)

	p.dream()

	p.setAge2(100)

	fmt.Printf("p = %#v\n", p)

	p1 := new(person)

	p1.name = "bbb"
	p1.age = 23
	fmt.Printf("p = %#v\n", p1)
	p1.setAge(100)
	fmt.Printf("p = %#v\n", p1)

	var a MyInt = 1

	a.hello()
}

/*方法和接受者
Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。

方法的定义格式如下：

func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
其中，

接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
方法名、参数列表、返回参数：具体格式与函数定义相同。
*/
func (p person) dream() {
	fmt.Printf("%s的梦想是学好Go语言", p.name)
}

/*指针类型的接收者
指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，
在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的this或者self。
例如我们为Person添加一个SetAge方法，来修改实例变量的年龄。
*/
func (p *person) setAge(age int) {
	p.age = age
}

/*值类型的接收者
当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
*/
func (p person) setAge2(age int) {
	p.age = age
}

/*
什么时候应该使用指针类型接收者
需要修改接收者中的值
接收者是拷贝代价比较大的大对象
保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/

/*任意类型添加方法
在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
*/
type MyInt int

func (m MyInt) hello() {
	fmt.Printf("Hello, 我是一个int。我的值是%d\n", m)
}

// 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
// func (i int) hello() {//不可以报错
// 	fmt.Printf("Hello, 我是一个int。我的值是%d", i)
// }

/*结构体的匿名字段
结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，
结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。


*/
//Person 结构体Person类型
type Person struct {
	string
	int
}

func noNameFieldStruct() {

	p := Person{
		"茂茂", 20,
	}

	fmt.Printf("%#v\n", p)
	fmt.Println(p.string, p.int)
}

//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
}

/*嵌套结构体
一个结构体中可以嵌套包含另一个结构体或结构体指针，就像下面的示例代码那样。
*/
func nestedDomo() {

	user := User{
		Name:   "韩梅梅",
		Gender: "女",
		Address: Address{
			Province: "广东",
			City:     "深圳",
		},
	}

	fmt.Printf("%#v\n", user)

	var user2 User

	user2.Address.Province = "北京"
	user2.City = "西城区" // 匿名字段可以省略

	fmt.Printf("user2=%#v\n", user2)
}

/*结构体的“继承”
Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
*/

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) bark() {
	fmt.Printf("%s会汪汪叫\n", d.name)
}

/*模仿继承 假继承
 */
func imitateExtends() {

	var d Dog

	d.Animal = &Animal{name: "Tim"}

	d.move() // d.Animal.move()

	d.bark()
}

/* 结构体字段的可见性
结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
*/
