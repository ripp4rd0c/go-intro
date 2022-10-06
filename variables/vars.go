package main

import "fmt"

var x = "hello world"

func main() {
	var (
		x string = "hello, world"

		y = "bla"

		z = "also bla"
	)
	x += "heheheheh"

	fmt.Println(x)

	fmt.Println(y == z)

	fmt.Println("end of program and we didnt use var h")
	f()

	const (
		a = 2
		b = 3
		c = 4
	)

	var (
		m int
	)

	fmt.Println("enter number: ")
	fmt.Scanf("%d", &m)
	fmt.Println(m*a, m*b, m*c)
}

func f() {
	fmt.Println(x)
}
