package main

import "fmt"

func main() {
	//there is only for loop in Golang , no while or do while loops

	fmt.Println("*********Simple For Loop************")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("*********while- For Loop************")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("*********Infinite - For Loop************")
	// need to use break statement to come out of infinite loop
	k := 1
	for {
		k = k + 1
		fmt.Println(k)
		if k > 5 {
			break
		}
	}

	fmt.Println("*********For-Range Loop************")
	//used with arrays, sclice maps and structs

	arrays := []int{1, 2, 3, 4, 50}

	for i, v := range arrays {
		fmt.Println(i, "-", v)
	}

	fmt.Println("*********For-Range Loop for only values************")
	for _, v := range arrays {
		fmt.Println(v)
	}
}
