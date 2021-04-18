package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	mapDemo8()
}

/*map定义
map[KeyType]ValueType
其中，

KeyType:表示键的类型。
ValueType:表示键对应的值的类型。
map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

make(map[KeyType]ValueType, [cap])
其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。


*/
func mapDemo1() {
	scoreMap := make(map[string]int, 4)

	fmt.Println(scoreMap) //map[]

	scoreMap["Tim"] = 100

	scoreMap["Tom"] = 90

	fmt.Println(scoreMap)                             //map[Tim:100 Tom:90]
	fmt.Println(scoreMap["Tom"])                      //90
	fmt.Printf("type of scoreMap is %T \n", scoreMap) //type of scoreMap is map[string]int
	fmt.Println(scoreMap["jimmy"])                    //0  不存在是返回默认值
}

/*map也支持在声明的时候填充元素
 */
func mapDemo2() {

	userInfo := map[int]string{
		1: "张三",
		2: "李四",
	}
	fmt.Println(userInfo)
}

/*判断某个键是否存在
 */
func mapDemo3() {

	userInfo := map[int]string{
		1: "张三",
		2: "李四",
	}

	value, ok := userInfo[0]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在此人")
	}
}

/*map 遍历
注意： 遍历map时的元素顺序与添加键值对的顺序无关
*/
func mapDemo4() {

	uInfo := make(map[int]string)
	uInfo[1] = "张西西"
	uInfo[2] = "李三炮"
	uInfo[3] = "万瑶瑶"
	uInfo[4] = "x笑笑"
	uInfo[5] = "ax"
	for k, v := range uInfo {
		fmt.Printf("%v:%v\n", k, v)
	}

	//但我们只想遍历key的时候，可以按下面的写法：
	for k := range uInfo {
		fmt.Printf("%v\n", k)
	}
}

/*使用delete()函数删除键值对
使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：

delete(map, key)
其中，

map:表示要删除键值对的map
key:表示要删除的键值对的键

*/
func mapDemo5() {

	uInfo := map[int]string{
		1: "xxx",
		2: "ccc",
	}
	fmt.Printf("%v - %p\n", uInfo, &uInfo) //map[1:xxx 2:ccc] - 0xc000126018
	delete(uInfo, 1)

	fmt.Printf("%v - %p\n", uInfo, &uInfo) //map[2:ccc] - 0xc000126018
}

/*按照指定顺序遍历map
 */
func mapDemo6() {
	rand.Seed(time.Now().UnixNano())

}

/*元素为map类型的切片
 */
func mapDemo7() {

	var mapSlice = make([]map[string]string, 3)

	for i, v := range mapSlice {
		fmt.Printf("i = %v,v = %v\n", i, v)
	}
	fmt.Println("after init value")

	mapSlice[0] = map[string]string{"哈哈": "笑"}
	mapSlice[1] = make(map[string]string, 10)
	mapSlice[1]["aa"] = "aaa"

	mapSlice[1]["bb"] = "bbb"

	mapSlice[1]["cc"] = "ccc"

	mapSlice[2] = map[string]string{"0001": "李华"}

	for i, v := range mapSlice {
		fmt.Printf("i = %v,v = %v\n", i, v)
	}
}

/*值为切片类型的map

 */
func mapDemo8() {
	var mapSlice = make(map[string][]string, 3)
	fmt.Println(mapSlice)
	fmt.Println("after init value")

	key := "广东省"

	value, ok := mapSlice[key]

	if !ok {
		value = make([]string, 0, 2)
	}

	value = append(value, "广州", "深圳")
	mapSlice[key] = value
	fmt.Println(mapSlice)

}
