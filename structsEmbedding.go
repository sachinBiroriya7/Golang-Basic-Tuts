package main

import "fmt"

//Instead of an inheritance model. GO uses a model that's similar to inheritance called composition
//it supports composition through embedding.

type person struct {
	name    string
	age     int
	address []string
}

type emp struct {
	/*
		name string
		age int
		address []string
	*/

	//since name age and address properties are already used in previous struct we can embedd it to use in out current struct
	person //embedding person inside emp struct for reusing the code which is already
	emp_id int
	salary int
}

func main() {
	person1 := person{
		name: "sachin",
		age:  23,
		address: []string{
			"LPU",
			"Punjab",
		},
	}

	emp1 := emp{}
	emp1.name = "hima" // using property of person struct through emp struct
	emp1.emp_id = 123  //as person is embedded in emp struct
	emp1.salary = 320012
	emp1.age = 21
	emp1.address = []string{"DU", "delhi"}

	fmt.Println(person1)
	fmt.Println(emp1)
}

/*
output :
{sachin 23 [LPU Punjab]}
{{hima 21 [DU delhi]} 123 320012}

*/

/*****************************************************************************
So it's not a traditional inheritance relationship where a emp is a person. Instead, it's a composition relationship, which is answering the question has a, so a emp has a person, or has person like characteristics. But it is not the same thing as an person, they cannot be used interchangeably .


In order to use data interchangeably, we're going to have to use interfaces
/*****************************************************************************/
