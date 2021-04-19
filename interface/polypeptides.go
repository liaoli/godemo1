package main

import "fmt"

//多肽
func main() {

	test(dog{name: "vvv"})
	test(car{brand: "porsche"})
}

type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	brand string
}

func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func test(m Mover) {
	m.move()
}
