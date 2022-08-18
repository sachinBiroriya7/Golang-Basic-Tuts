package main

import "fmt"

//creating a struct

type employee struct {
	name    string
	age     int
	id      int
	salary  float32
	address []string
}

func main() {

	// assigning values

	emp1 := employee{
		name:   "sachin",
		age:    23,
		id:     4349,
		salary: 35000.50,
		address: []string{
			"dhungadhara", "almora",
		},
	}

	fmt.Println(emp1)
	fmt.Println(emp1.salary) //accessing values within struct
	fmt.Println(emp1.address[1])

	//anonymous struct -> struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code
	tempEmp := struct {
		name string
		age  int
	}{name: "surya", age: 23}
	fmt.Println(tempEmp)
	//To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type

	// Structs are pass by values
	emp2 := emp1
	emp2.name = "himakshi"
	emp2.id = 2313
	fmt.Println(emp2)
	fmt.Println(emp1)

	// to change uderlying data we need to use pointers
	emp3 := &emp2
	emp3.name = "Hima"
	emp3.age = 21

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
