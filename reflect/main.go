package main

import (
	"fmt"
	"reflect"
)

func main() {

	nilAndValid()
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type :%v kind : %v\n", v.Name(), v.Kind())
}

type persion struct {
	name string
	age  int
}

type book struct {
	name  string
	price float32
}

func typeOfDemo() {
	var a float32 = 3.14
	reflectType(a)

	var b int64 = 1
	reflectType(b)
	var p *int64 = new(int64)
	*p = 100
	reflectType(p)

	pe := persion{
		"dd",
		22,
	}

	bk := book{
		"金瓶梅",
		55.79,
	}

	reflectType(pe)

	reflectType(bk)
}

// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。

/*
reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。
reflect.Value与原始值之间可以互相转换。
*/
func reflectValue(x interface{}) {

	v := reflect.ValueOf(x)
	k := v.Kind()

	switch k {
	case reflect.Int32:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int32(v.Int()))
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	case reflect.Bool:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %v\n", bool(v.Bool()))

	}

}

func valueOfDemo() {
	var a float32 = 3.14
	var b int64 = 100
	var c int32 = 111
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
	reflectValue(true)
}

/*通过反射设置变量的值
想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。
而反射中使用专有的Elem()方法来获取指针对应的值。
*/

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func demo() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	fmt.Println(a)
	reflectSetValue2(&a)
	fmt.Println(a)
}

/*
isNil()和isValid()
isNil()
func (v Value) IsNil() bool
IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

isValid()
func (v Value) IsValid() bool
IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
*/

func nilAndValid() {
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())

}
