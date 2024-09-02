package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var numberA, numberB, numberC int

	_, err := fmt.Scan(&numberA, &numberB, &numberC)
	if err != nil {
		log.Fatal(err)
	}

	mySlice := []int{numberA, numberB, numberC}

	sort.Ints(mySlice)

	if isTriangle(mySlice) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isTriangle(mySlice []int) bool {
	if mySlice[0] + mySlice[1] > mySlice[2] {
		return true
	} else {
		return false
	}
}
