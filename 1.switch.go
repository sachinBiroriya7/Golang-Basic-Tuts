package main

import "fmt"

// break statements in swirch cases are implicitly added so we do not need to add, unless needed
// we do not need to add brackets {} in cases as well

//type 1: simple switch case
func main() {
	var i int
	fmt.Println("enter value ")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	case 3:
		fmt.Println("case 3")
	case 4:
		fmt.Println("case 4")
	default:
		fmt.Println("case not found")
	}
	fmt.Println("Switch type 1 ended")
	fmt.Println("********************************")

	//type 2: switch for declaring and performing operation

	switch j := 10 + i; j {
	case 11:
		fmt.Println("case 11")
	case 12:
		fmt.Println("case 12")
	case 13:
		fmt.Println("case 13")
	case 14:
		fmt.Println("case 14")
	case 15, 16, 17: // one output match will run this
		fmt.Println("case in 15,16 or 17")
	default:
		fmt.Println("case not found")
	}
	fmt.Println("switch type 2 ended")
	fmt.Println("********************************")

	//Type 3: Type switch
	//A type switch compares types instead of values. You can use this to discover the type of an interface value.

	var k interface{} = "sachin"
	//var k interface{} = 10.1
	//var k interface{} = []int{}
	//var k interface{} = 1

	switch k.(type) {
	case int:
		fmt.Println("integer type")
	case float64:
		fmt.Println("Float type")
	case string:
		fmt.Println("string type")
	default:
		fmt.Println("another Type")
	}

}
