package main

import (
	"fmt"
	"strconv"
)

var wiek map[string]uint8 = make(map[string]uint8)

func main() {

	wiek["Adam"] = 26
	wiek["Aga"] = 23

	check("Aga")
	check("Agaa")

	fmt.Println("\nList: ")
	for name, age := range wiek {
		fmt.Println(name + "'s age: " + strconv.Itoa(int(age)))
	}
}

func check(value string) {
	var age, ok = wiek[value]
	if ok {
		fmt.Println(value + "'s age is " + strconv.Itoa(int(age)))
	} else {
		fmt.Println("Error " + value + " don't exist in map!")
	}
}
