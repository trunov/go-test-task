package solutionOne

import (
	"sort"
)

func sumArr(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
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
				break
			}

			m[s] = []int{C[i]}
			// redundant
			break
		}

		if i != 0 && i < len(S)-1 {
			if s == string(S[i+1]) || s == string(S[i-1]) {
				v, ok := m[s]
				if ok {
					m[s] = append(v, C[i])
					continue
				}

				m[s] = []int{C[i]}
			}
		}

	}

	for k, v := range m {
		sort.Ints(m[k])
		sum += sumArr(v[:len(v)-1])
	}

	return sum
}
