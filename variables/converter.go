package main

import "fmt"

func main() {

	var (
		F    int
		FEET float64
	)

	fmt.Println("fahrenheit to celsius converter\nenter f: ")

	fmt.Scan(&F)

	fmt.Println("in celsius ", (F-32)*5/9)

	fmt.Println("feet to meter converter")

	fmt.Scan(&FEET)

	fmt.Println("in meter: ", FEET*0.3048)
}
