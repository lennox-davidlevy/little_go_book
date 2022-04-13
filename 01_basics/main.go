package main

import (
	"fmt"
)

func main() {
	name, power := "Goku", getPower()
	fmt.Printf("%s's power is over %d\n", name, power)
}

func getPower() int {
	return 9001
}

func log(message string) {

}

func add(a int, b int) int {
	return a + b
}
