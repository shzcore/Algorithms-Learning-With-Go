package main

import (
	"Algorithms-Learning-With-Go/sort/utils"
	"log"

	"Algorithms-Learning-With-Go/sort/quick"
	// "go-algorithms/sort/bubble"
)

func main() {
	list := utils.GetArrayOfSize(10)
	log.Println(list)
	quick.HoareSort(list, 0, len(list)-1)
	// quick.LomutoSort(list, 0, len(list)-1)
	// quick.Sort(list)
	// bubble.Sort(list)
	log.Println(list)
}
