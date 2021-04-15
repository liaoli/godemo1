package main

import "fmt"

func main() {
	// fmt.Println("Hello World")

	// fmt.Print("a", "b", "c")

	// fmt.Println()

	// fmt.Println("a", "b", "c")

	// fmt.Printf("hello world")

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
