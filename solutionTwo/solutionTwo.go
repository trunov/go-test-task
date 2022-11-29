package solutionTwo

import (
	"fmt"
	"log"
	"strconv"
)

func findMax(a []int) (max int) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func Solution(A []int) int {
	var maxSubset int
	maxNum := fmt.Sprintf("%b", findMax(A))
	result := make([][]int, len(A))

	for i, v := range A {
		num := 1 << 24
		s := fmt.Sprintf("%b", num|v)

		for j := 0; j < len(maxNum); j++ {
			n, err := strconv.Atoi(string(s[len(s)-j-1]))
			if err != nil {
				log.Fatal("Conversion of string to int did not work")
			}

			result[i] = append(result[i], n)
		}
	}

	for j := 0; j < len(maxNum); j++ {
		colsum := 0
		for i := 0; i < len(A); i++ {
			colsum = colsum + result[i][j]
		}

		if colsum > maxSubset {
			maxSubset = colsum
		}
	}

	return maxSubset
}
