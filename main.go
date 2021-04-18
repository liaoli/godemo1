package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func main() {
	// fmt.Println("Hello World")

	// fmt.Print("a", "b", "c")

	// fmt.Println()

	// fmt.Println("a", "b", "c")

	// fmt.Printf("hello world")

	//ConstantDeclarationAndInitialization()
	// floatType()

	// TypeDefaultValue()
	// stringT

	// byteType()

	// typeConversion()
	// other2String()
	string2Num()

}

func string2Num() {
	//1、string类型转换成整型

	str := "123456"
	fmt.Printf("%v--%T\n", str, str)

	/*
		ParseInt
		参数1：string数据
		参数2：进制
		参数3：位数 32 64 16

	*/
	num, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%v--%T", num, num)

	/*
		ParseFloat
		参数1：string数据
		参数2：位数 32 64

	*/
	str1 := "123456.333xxxx"
	num1, _ := strconv.ParseFloat(str1, 64)
	fmt.Printf("%v--%T", num1, num1)

	//不建议把string类型转换成bool型
	b, _ := strconv.ParseBool("xxxxxxx") // string 转 bool
	fmt.Printf("值：%v 类型：%T", b, b)
}

func other2String() {
	//1、通fmt.Sprintf() 把其他类型转换成 String 类型

	//注意：Sprintf 使用中需要注意转换的格式 int 为%d  float 为%f  bool 为%t   byte 为%c

	/*
		var i int = 20
		var f float64 = 12.456
		var t bool = true
		var b byte = 'a'

		str1 := fmt.Sprintf("%d", i)
		fmt.Printf("值：%v 类型：%T\n", str1, str1)

		str2 := fmt.Sprintf("%.2f", f)
		fmt.Printf("值：%v 类型：%T\n", str2, str2)

		str3 := fmt.Sprintf("%t", t)
		fmt.Printf("值：%v 类型：%T\n", str3, str3)

		str4 := fmt.Sprintf("%c", b)
		fmt.Printf("值：%v 类型：%T\n", str4, str4)
	*/

	//2、通过strconv  把其他类型转换成string类型

	/*
		FormatInt
		参数1：int64 的数值
		参数2：传值int类型的进制
	*/
	var i int = 20
	str1 := strconv.FormatInt(int64(i), 10)
	fmt.Printf("值：%v 类型：%T\n", str1, str1)

	/*
		参数 1：要转换的值
		参数 2：格式化类型 'f'（-ddd.dddd）、
			 'b'（-ddddp±ddd，指数为二进制）、
			 'e'（-d.dddde±dd，十进制指数）、
			 'E'（-d.ddddE±dd，十进制指数）、
			 'g'（指数很大时用'e'格式，否则'f'格式）、
			 'G'（指数很大时用'E'格式，否则'f'格式）。
		 参数 3: 保留的小数点 -1（不对小数点格式化）
		 参数 4：格式化的类型 传入 64  32
	*/
	var f float32 = 20.231313
	str2 := strconv.FormatFloat(float64(f), 'f', 4, 32)
	fmt.Printf("值：%v 类型：%T\n", str2, str2)

	str3 := strconv.FormatBool(true) //没有任何意义
	fmt.Printf("值：%v 类型：%T\n", str3, str3)

	a := 'b' //没有任何意义
	str4 := strconv.FormatUint(uint64(a), 10)
	fmt.Printf("值：%v 类型：%T\n", str4, str4) //值：98 类型：string
}

func typeConversion() {
	//1、整型和整型之间的转换
	var a int8 = 20
	var b int16 = 40
	fmt.Println(int16(a) + b)

	//2、浮点型和浮点型之间的转换
	var f1 float32 = 20
	var f2 float64 = 40
	fmt.Println(float64(f1) + f2)

	//3、整型和浮点型之间的转换
	var a1 float32 = 20.23
	var b1 int = 40
	fmt.Println(a1 + float32(b1))

	//注意：转换的时候建议从低位转换成高位，高位转换成低位的时候如果转换不成功就会溢出，和我们想的结果不一样。

}

func byteType() {
	//1、golang中定义字符  字符属于int类型

	var c = 'a'
	fmt.Printf("flag = %v type is %T  size is %v\n", c, c, unsafe.Sizeof(c))
	//2、原样输出字符
	fmt.Printf("flag = %c type is %T  size is %v\n", c, c, unsafe.Sizeof(c))
	//3、定义一个字符串输出字符串里面的字符
	str := "dfsdf"

	fmt.Printf("value = %v flag = %c type is %T  size is %v \n", str[2], str[2], str[2], unsafe.Sizeof(str[2]))
	//4、一个汉字占用 3个字节(utf-8), 一个字母占用一个字节
	// unsafe.Sizeof() 没法查看string类型数据所占用的存储空间

	var str1 = "this" //占用4个字节
	fmt.Println(len(str1))

	var str2 = "你好go" //8
	fmt.Println(len(str2))

	var char = '国'

	fmt.Printf("value = %v flag = %c type is %T  size is %v \n", char, char, char, unsafe.Sizeof(char))

	//6、通过循环输出字符串里面的字符
	s := "你好 golang"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("value = %v flag = %c type is %T  size is %v \n", s[i], s[i], s[i], unsafe.Sizeof(s[i]))
	}

	// s := "你好 golang"
	for _, v := range s { //rune
		fmt.Printf("value = %v flag = %c type is %T  size is %v \n", v, v, v, unsafe.Sizeof(v))
	}

	//7、修改字符串
	s1 := "big"
	byteStr := []byte(s1)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr))

	s2 := "你好golang"
	runeStr := []rune(s2)
	runeStr[0] = '大'
	fmt.Println(string(runeStr))

}

func printChar(a string) {
	fmt.Printf("flag = %v type is %T  size is %v\n", a, a, unsafe.Sizeof(a))

}

func stringType() {
	//1、定义string类型
	var str1 string = "你好 golang"

	var str2 = "你好 go"

	str3 := "你好gogogo"

	print(str1)
	print(str2)
	print(str3)
	//2、字符串转义符
	str4 := "this is \n hahaha" //换行
	print(str4)

	str5 := "C:\\Go\\bin" //C:\Go\bin    输出反斜杠
	fmt.Println(str5)

	/*
		C:Go"bin"
	*/
	str6 := "C:Go\"bin\"" //输出双引号
	fmt.Println(str6)

	// 3、多行字符串   `(反引号)  tab键上面
	// 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

	str7 := `注意：
	1.布尔类型变量的默认值为false。
	2.Go 语言中不允许将整型强制转换为布尔型.
	3.布尔型无法参与数值运算，也无法与其他类型进行转换。`

	fmt.Println(str7)

	// 4、len(str) 求长度

	len := len(str1)

	fmt.Printf("lenght of len is %v\n", len)
	// 5、+ 或者 fmt.Sprintf拼接字符串

	str8 := str2 + str1

	str9 := fmt.Sprintf("%v%v", str1, str2)

	fmt.Println(str8)

	fmt.Println(str9)

	str1 += str2

	fmt.Println(str1)

	str10 := "反引号间换行将被作为字符串中的换" +
		"文本将会原样输出" +
		"行，但是所有的转义字符均无效"
	fmt.Println(str10)
	//6、strings.Split 分割字符串  strings需要引入strings包
	var str11 = "123-456-789"

	arr1 := strings.Split(str11, "-")

	fmt.Println(arr1)

	//7、strings.Join(a[]string, sep string) join 操作  表示把切片链接成字符串

	arr2 := []string{"I", "will", "always", "love", "you"}
	str12 := strings.Join(arr2, " ")

	fmt.Println(str12)

	//8、 strings.contains 判断是否包含

	contain1 := strings.Contains("hahaha", "haha")

	fmt.Println(contain1)

	contain2 := strings.Contains("hahaha", "hha")

	fmt.Println(contain2)

	//9、strings.HasPrefix,strings.HasSuffix 前缀/后缀判断

	hasPrefix := strings.HasPrefix("i am happy", "i")

	fmt.Println(hasPrefix)

	hasSuffix := strings.HasSuffix("i am happy", "y")

	fmt.Println(hasSuffix)

	//10、 strings.Index(),strings.LastIndex() 子串出现的位置  查找不到返回-1 查找到返回下标位置 下标是从0开始的

	index := strings.Index("hello golang", "llo")

	fmt.Println("index = ", index)

	index1 := strings.Index("hello golang", "lll")

	fmt.Println("index1 = ", index1)

	lastIndex := strings.LastIndex("hello golang", "lll")

	fmt.Println("lastIndex = ", lastIndex)

	lastIndex1 := strings.LastIndex("hello golang", "l")

	fmt.Println("lastIndex1 = ", lastIndex1)

}

func print(a string) {
	fmt.Printf("flag = %v type is %T  size is %v\n", a, a, unsafe.Sizeof(a))

}

func TypeDefaultValue() {

	/*
		Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

		注意：
		1.布尔类型变量的默认值为false。
		2.Go 语言中不允许将整型强制转换为布尔型.
		3.布尔型无法参与数值运算，也无法与其他类型进行转换。
	*/
	var flag bool
	//1.布尔类型变量的默认值为false。
	fmt.Printf("flag = %v type is %T  size is %v\n", flag, flag, unsafe.Sizeof(flag))

	// 2.string型变量的默认值为空。

	var s string
	fmt.Printf("%v\n", s)

	//3.int型变量的默认值为0。

	var i int
	fmt.Printf("int value is %v and type %T \n", i, i)

	//4.float型变量的默认值为0。

	var f float32
	fmt.Printf("f value is %v and type %T \n", f, f)

	//5、Go 语言中不允许将整型强制转换为布尔型.
	// var a = 1
	// if a { //错误写法
	// 	fmt.Printf("true")
	// }

	//6.布尔型无法参与数值运算，也无法与其他类型进行转换。

	// var s = "this is str"
	// if s { //错误写法
	// 	fmt.Printf("true")
	// }

	var f1 = false
	if f1 { //正确写法
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}

}

func floatType() {

	var a float32 = 3.12
	fmt.Printf("a = %f size of type %T is %v\n", a, a, unsafe.Sizeof(a))

	var f float64 = 3.12
	fmt.Printf("f = %f size of type %T is %v\n", f, f, unsafe.Sizeof(f))

	var pi = 3.1415926 //64位的系统中 浮点数默认是 float64

	fmt.Printf("pi = %v \n", pi) //%v 原样输出

	fmt.Printf("pi = %f \n", pi) //%f 输出float 类型数据

	fmt.Printf("pi = %.3f \n", pi) //%。3f 输出float 类型数据并指定保留几位小数

	fmt.Printf("size of %T is %v \n", pi, unsafe.Sizeof(pi))

	//科学计数法
	var f2 = 3.14e4

	fmt.Printf("%v -- %T\n", f2, f2)

	var f3 = 3.14e-2

	fmt.Printf("%v -- %T\n", f3, f3)

	// 5、Golang 中 float 精度丢失问题
	var f4 = 1129.6

	fmt.Printf("f4 = %v\n", f4*100)

	m1, m2 := 8.2, 3.8

	fmt.Printf("(m1 - m2) = %v\n", (m1 - m2))

	//6、int类型转换成float类型

	integer := 22
	ff := float64(integer)

	fmt.Printf("integer = %v  type is %T \n", integer, integer)

	fmt.Printf("integer = %v  type is %T \n", ff, ff)

	// 7、float类型转换成int类型 （不建议这样做）

	var f5 = 33.3

	i1 := int(f5)

	fmt.Printf("integer = %v  type is %T \n", f5, f5)

	fmt.Printf("integer = %v  type is %T  size is %v \n", i1, i1, unsafe.Sizeof(i1))

}

func intType() {

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
