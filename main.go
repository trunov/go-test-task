package main

import (
	"fmt"

	"github.com/trunov/go-test-task/solutionOne"
	"github.com/trunov/go-test-task/solutionTwo"
)

func main() {
	fmt.Println("Solution One")
	fmt.Println(solutionOne.Solution("abccd", []int{0, 1, 2, 3, 4}))
	fmt.Println(solutionOne.Solution("aabbcc", []int{1, 2, 1, 2, 1, 2}))
	fmt.Println(solutionOne.Solution("aaaa", []int{1, 2, 3, 4}))

	fmt.Println("Solution Two")
	fmt.Println(solutionTwo.Solution([]int{13, 7, 2, 8, 3}))
	fmt.Println(solutionTwo.Solution([]int{1, 2, 4, 8}))
	fmt.Println(solutionTwo.Solution([]int{16, 16}))
}
