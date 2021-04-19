package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := Person{
		name: "杨哈哈",
		age:  33,
	}

	data := []string{"吃饭", "睡觉", "打牌"}

	p1.setDreams(data)
	data[1] = "运动"
	fmt.Println(p1.dreams)

	p1.setDreams2(data)
	data[1] = "泡妞"
	fmt.Println(p1.dreams)
	//同样的问题也存在于返回值slice和map的情况，在实际编码过程中一定要注意这个问题。

}

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

/*结构体标签（Tag）
Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：

`key1:"value1" key2:"value2"`
结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。
同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。

注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。
结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，
通过反射也无法正确取值。例如不要在key和value之间添加空格。
*/

func jsonDemo1() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu%02d", i),
		}
		c.Students = append(c.Students, stu)
	}

	fmt.Printf("%#v", c)

	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)

	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	fmt.Printf("json:%s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体

	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`

	c1 := &Class{}

	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json Unmarshal failed")
		return
	}

	fmt.Printf("%#v", c1)
}

//
type Employee struct {
	ID     int    `json:"id"`     //通过指定tag实现json序列化该字段时的key
	Gender string `json:"gender"` //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func jsonDemo2() {
	s := Employee{
		ID:     1,
		Gender: "女",
		name:   "颜色",
	}

	data, err := json.Marshal(s)

	if err != nil {
		fmt.Println("json Unmarshal failed")
		return
	}

	fmt.Printf("json:%s\n", data)

}

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) setDreams(dreams []string) {
	p.dreams = dreams
}

func (p *Person) setDreams2(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
