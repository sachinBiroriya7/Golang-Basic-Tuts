/*
in Go, we don't have exceptions, because a lot of cases that are considered exceptional in other languages are considered normal in a go application. For example, if you try and open a file, and that file doesn't exist, that's actually a pretty normal response.

Errors are returned to handle such situation and this resposiblity is on devs

However, there are some things that get a go application into a situation where it cannot continue and that is considered exceptional. But instead of using the word exception, we're going to use another word and that word is going to be panic
*/

//panic -> cause our app to panic and stop working further execution, we can use panic keyword in are Go program to make Go program in panic anything below panic wont be executed, even defer statements

package main

import "fmt"

func main() {

	a, b := 10, 0
	var result int
	//result = a / b // runtime exception -> program panics and stop further execution

	fmt.Println(result, a, b)
	fmt.Println("end")

	// Panics with defer statements

	fmt.Println("hello")

	fmt.Println("opening resource")
	defer fmt.Println("closing resourse")

	panic("cant continue the program")

	fmt.Println("world")

}
