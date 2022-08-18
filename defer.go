/*
the defer keyword is used to delay the execution of a function or a statement until the nearby function returns. In simple words, defer will move the execution of the statement to the very end inside a function.

the way the defer keyword works in go, is that it actually executes any functions that are passed into it, after the function finishes its final statement, but before it actually returns.

Most common use is to associate the opening of a resource and the closing of the resource right next to each other. usually for purposes of cleanup
*/

/*when there are multiple defer statements/funcs they are executed in  LIFO order or last in first out. because often we're going to use the deferred keyword to close out resources. And it makes sense that we close resources out in the opposite order that we open them, because one resource might actually be dependent on another one.
 */

package main

import "fmt"

func main() {
	defer fmt.Println("hello")

	fmt.Println("world")
	deferFunc2()

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	deferFunc()

}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func deferFunc2() {

	a := "start"
	defer fmt.Println("a =", a)
	a = "end"
	fmt.Println()

	/*. you might also realize the defer statement is going to cause this statement to print after the main function is done. And before main is done, we've already changed the value to end.

	So the output must be : end?   No, output will be start,  because ,

	when you defer a function like this, it actually takes the argument at the time the defer is called, not at the time the called function is executed. So even though the value of a is changed, before we leave the main function, we are going to actually eagerly evaluate this variable. And so the value start is going to be put in here,
	*/
}

/*
OUTPUT :

world
4
3
2
1
0
three
two
one
hello

*/
/*
flow of the above program

world,
deferFunc2(),
deferFunc() -> 0,1,2,3,4 will be stored in stack as print has defer keyword hence will be retured in LIFO order i.e -> 4,3,2,1,0
then in main function -> all defer statements in LIFO order

*/
