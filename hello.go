package main

import (
	"fmt"
)

func printText(texto string) {

	x := 10000
	var i int
	for i < x {
		i++
		fmt.Printf("%v\n", i ,texto)
	}

}

func main() {

	printText(".-ya me aburri XD \n")
}
