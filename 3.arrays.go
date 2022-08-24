package main

import "fmt"

func main() {
	fmt.Println("*******Array Declaration************")

	//four sysntaxs for array creation
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)

	var arr2 [2]string
	arr2[0] = "sachin"
	arr2[1] = "singh"
	fmt.Println(arr2)

	arr3 := [3]int{11, 12, 13}
	fmt.Println(arr3)

	arr4 := []string{"sachin", "singh", "biroriya", "GoLang", "COde"}
	fmt.Println(arr4)

	fmt.Println("*******Array Traversal************")
	//using len() Function
	fmt.Println("arr1 length :", len(arr))
	fmt.Println("arr2 length :", len(arr2))
	fmt.Println("arr3 length :", len(arr3))
	fmt.Println("arr4 length :", len(arr4))

	fmt.Println("*******Array Traversal************")
	//simple for loop
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], ",")
	}
	fmt.Println("-------")

	//for Range loop
	for index, value := range arr2 {
		fmt.Println(index, "-", value)
	}
}
