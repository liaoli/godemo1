package main

import "fmt"

func main() {
	// arrayDemo1()
	arrayTest()

	a := [3]int{10, 20, 30}

	b := [...]int{10, 20, 30}

	fmt.Println(a == b)
}

//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
func arrayTest() {
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]

	arr1 := [3]int{10, 20, 30}

	arr2 := [...]int{10, 20, 30}

	fmt.Println(arr1 == arr2)

}
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func twoDimensionalArray() {
	arr := [][]string{
		{
			"北京", "上海",
		},
		{
			"广州", "深圳",
		},
		{
			"武汉", "长沙",
		},
	}

	fmt.Println(arr[2][1])

	for index1, v1 := range arr {
		for index2, v2 := range v1 {
			fmt.Printf("(%v,%v) = %v\n", index1, index2, v2)
		}
	}
}

func traversal() {

	arr := []int{2, 5, 6, 8, 9}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v--%T \n", arr[i], arr[i])
	}

	for index, value := range arr {
		fmt.Printf("(%v,%v)\n", index, value)
	}

}

func arrayDemo1() {
	var testArray [3]int
	var numArray = [3]int{1, 2}
	var cityArray = []string{"北京", "上海", "广州", "深圳"}
	var cc = [...]int{1, 2, 4, 5}

	var arr = [...]string{2: "杭州", 7: "深圳"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)
	fmt.Println(cc)
	fmt.Println(arr)
	fmt.Println(len(arr))

}
