package main

import "fmt"

func main() {
	var item1, item2 string
	item2 = "Write Intro to Go (part 1)"

	var item3 string = "Listen to the Go Time podcast"

	item4 := "Read the Golang Weekly email"

	for i := 0; i < 5; i++ {
		fmt.Println("My TODOs:")
		if item1 == "" {
			fmt.Println("Empty TODO")
		} else if item1 == "todo" || item1 == "wip" {
			fmt.Println("TODO: Write this TODO")
		} else {
			fmt.Println(item1)
		}
		fmt.Println(item2)
		fmt.Println(item3)
		fmt.Println(item4)
		if i < 5 && i != 4 {
			fmt.Println()
		}
	}
}
