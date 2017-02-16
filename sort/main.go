package main

import (
	"go-algorithms/sort/utils"
	"log"

	"go-algorithms/sort/quick"
	// "go-algorithms/sort/bubble"
)

func main() {
	list := utils.GetArrayOfSize(10)
	log.Println(list)
	quick.LomutoSort(list, 0, len(list)-1)
	// quick.Sort(list)
	// bubble.Sort(list)
	log.Println(list)
}
