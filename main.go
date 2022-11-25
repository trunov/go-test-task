package main

import (
	"fmt"

	"github.com/trunov/go-test-task/solutionOne"
)

func main() {
	fmt.Println(solutionOne.Solution("abccd", []int{0, 1, 2, 3, 4}))
	fmt.Println(solutionOne.Solution("aabbcc", []int{1, 2, 1, 2, 1, 2}))
	fmt.Println(solutionOne.Solution("aaaa", []int{1, 2, 3, 4}))

}
