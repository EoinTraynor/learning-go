package main

import "fmt"

func main()  {
	sentence := "Look at my sentence"
	for index,  letter := range sentence {
		if index % 2 != 0 {
			fmt.Println(index, string(letter))
		}
	}
}
