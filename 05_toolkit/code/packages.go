package main

import (
	"fem-intro-to-go/05_toolkit/code/exercise_5a_solution/utils"
	"fmt"
)

func calculateData() int {
	totalData := utils.Add(1, 2, 3, 4, 5)
	return totalData
}

func main() {
	fmt.Println("Packages!")
	total := calculateData()
	fmt.Println(total)
}
