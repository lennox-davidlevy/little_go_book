package main

import "fmt"

//Initialize empty array with size of 10
// func main() {
// 	var scores [10]int
// 	scores[0] = 339
// 	fmt.Println(scores)
// }

//Initialize array with size and values
// func main() {
// 	scores := [4]int{9001, 9333, 212, 33}
// 	for index, value := range scores {
// 		fmt.Println(index, value)
// 	}
// }

//Rarely use arrays, instead Go uses slices
// func main() {
// 	scores := []int{1, 3, 293, 4, 9}
// 	for index, value := range scores {
// 		fmt.Println(index, value)
// 	}
// }

//Create slice with make
//func main() {
//	//make([]int, length, capacity)
//	scores := make([]int, 2, 10)
//	for index, value := range scores {
//		fmt.Println(index, value)
//	}
//}

// func main() {
// 	scores := make([]int, 0, 10)
// 	scores = append(scores, 5)
// 	fmt.Println(scores)
// }

// func main() {
// 	scores := make([]int, 0, 10)
// 	scores = scores[0:8]
// 	scores[7] = 9033
// 	fmt.Println(scores)
// }

// func main() {
// 	scores := make([]int, 0, 5)
// 	c := cap(scores)
// 	fmt.Println(c)

// 	for i := 0; i < 25; i++ {
// 		scores = append(scores, i)
// 	}

// 	// if our capacity has changed,
// 	// Go had to grow our array to
// 	// accomodate the new data

// 	if cap(scores) != c {
// 		c = cap(scores)
// 		fmt.Println(c)
// 	}
// }

// func main() {
// 	scores := make([]int, 5)
// 	scores = append(scores, 9332)
// 	fmt.Println(scores)
// }

// four common ways to initialize a slice
func main() {
	names := []string{"leto", "jessica", "paul"}
	checks := make([]bool, 10)
	var other_names []string
	scores := make([]int, 0, 20)
	fmt.Println(names, checks, other_names, scores)
}
