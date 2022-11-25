package solutionOne

import (
	"sort"
)

func calculateSum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Solution(S string, C []int) int {
	sum := 0
	m := make(map[string][]int, len(C))

	for i, v := range S {
		s := string(v)

		if i == 0 && s == string(S[i+1]) {
			m[s] = []int{C[i]}
			continue
		}

		if i == len(S)-1 && s == string(S[i-1]) {
			v, ok := m[s]
			if ok {
				m[s] = append(v, C[i])
				sort.Ints(m[s])
				break
			}

			m[s] = []int{C[i]}
		}

		if i != 0 && i < len(S)-1 {
			if s == string(S[i+1]) || s == string(S[i-1]) {
				v, ok := m[s]
				if ok {
					m[s] = append(v, C[i])
					sort.Ints(m[s])
					continue
				}

				m[s] = []int{C[i]}
			}
		}

	}

	for _, v := range m {
		if len(v) > 1 {
			sum += calculateSum(v[:len(v)-1])
		}
	}

	return sum
}
