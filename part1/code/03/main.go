package main

import "fmt"

func main() {
	var item1, item2 string
	item2 = "Write Intro to Go (part 1)"

	var item3 string = "Listen to the Go Time podcast"

	item4 := "Read the Golang Weekly email"

	for i := 0; i < 5; i++ {
		fmt.Println("My TODOs:")
		fmt.Println(item1)
		fmt.Println(item2)
		fmt.Println(item3)
		fmt.Println(item4)
		fmt.Println()
	}

	i := 0
	for i < 5 {
		fmt.Println("My TODOs:")
		fmt.Println(item1)
		fmt.Println(item2)
		fmt.Println(item3)
		fmt.Println(item4)
		fmt.Println()
		i++
	}
}
