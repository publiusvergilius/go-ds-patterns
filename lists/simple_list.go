package main

import (
	"container/list"
	"fmt"
)

/**
 * Linear list of heterogeneous elements 
 */

func simple_list () {
	var numList list.List

	// Add to the back
	numList.PushBack(7)
	numList.PushBack(8)
	numList.PushBack(6)

	numList.PushBack(6.4)

	numList.PushBack("The end!")

	for element := numList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

}
