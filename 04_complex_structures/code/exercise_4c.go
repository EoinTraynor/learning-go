package main

import "fmt"

func getAverage(scores [5]float64) float64 {
	total := 0.0
	for _, score := range scores {
		total += score
	}
	average := total / float64(len(scores))
	return average
}

var pets map[string]string = map[string]string{
	"nip":    "Dog",
	"django": "Dog",
	"Pus":    "Cat",
	"Jerry":  "Mouse",
}

func existsInPets(key string) bool {
	_, isPresent := pets[key]
	return isPresent
}

func addGroceriesToList(currentList []string, groceriesToAdd []string) []string {
	for _, item := range groceriesToAdd {
		currentList = append(currentList, item)
	}
	return currentList
}

func main() {
	var myScores [5]float64
	myScores[0] = 1.1
	myScores[1] = 1.21
	myScores[2] = 7.1
	myScores[3] = 1.6
	myScores[4] = 4.1

	var groceries = []string{
		"apple", "bannan", "pear", "kiwi", "orange",
	}
	var extraGroceries = []string{
		"grapes", "berries",
	}

	average := getAverage(myScores)
	doesNipExist := existsInPets("nip")
	fullList := addGroceriesToList(groceries, extraGroceries)
	fmt.Println(average)
	fmt.Println(doesNipExist)
	fmt.Println(fullList)
}
