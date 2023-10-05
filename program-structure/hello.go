package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	/*
	 * This is my first sample program.
	 */
	var countryCapitalMap map[string]string
	/* create a map*/
	countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map*/
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* test if entry is present in the map or not*/
	capital, ok := countryCapitalMap["United States"]

	/* if ok is true, entry is present otherwise entry is absent*/
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}
	var x float64
	x = 20
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	var greeting = "hello world"
	for i := 0; i < len(greeting); i++ {
		fmt.Println("%x ", greeting[i])
	}
	area = max(LENGTH, WIDTH)
	fmt.Printf("value of area : %d\n", area)
	fmt.Println(x)
	fmt.Println("x is of type %T\n", x)

	var aa int = 20 /* actual variable declaration */
	var ip *int     /* pointer variable declaration */

	ip = &aa /* store address of aa in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &aa)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
