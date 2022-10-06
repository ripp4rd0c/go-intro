package main

import "fmt"

func main() {

	fmt.Println("this section is for types")
	var a int32 = 43434
	var c rune = 4444

	//var d float64 = 44.444

	var str string = `hello, world`
	fmt.Println(len(str))
	fmt.Println(string(str[1]))
	fmt.Println(len("HELLO"))
	fmt.Println("blabla bitch")
	fmt.Println(a, c, str)
	fmt.Println(string((byte(98))))

	// booleans

	fmt.Println(true && true)

	fmt.Println(true || false)
	fmt.Println(!false)

	fmt.Println(32132 * 42452)
}
