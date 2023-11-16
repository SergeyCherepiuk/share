package lcs

import (
	"slices"
)

func Diff(linesPrev, linesCurr [][]byte) (deletions, insertions []int) {
	deletions = make([]int, 0)
	insertions = make([]int, 0)

	length := length(linesPrev, linesCurr)

	i, j := len(linesPrev), len(linesCurr)
	for i != 0 && j != 0 {
		if slices.Equal(linesPrev[i-1], linesCurr[j-1]) {
			i, j = i-1, j-1
			continue
		}

		if length[i-1][j] <= length[i][j-1] {
			j--
			insertions = append(insertions, j)
		} else {
			i--
			deletions = append(deletions, i)
		}
	}

	for ; i > 0; i-- {
		deletions = append(deletions, i-1)
	}

	for ; j > 0; j-- {
		insertions = append(insertions, j-1)
	}

	slices.Reverse(insertions)
	slices.Reverse(deletions)
	return
}

func length(linesPrev, linesCurr [][]byte) [][]int {
	results := make([][]int, len(linesPrev)+1)
	for i := range results {
		results[i] = make([]int, len(linesCurr)+1)
	}

	for i, result := range results {
		for j := range result {
			if i == 0 || j == 0 {
				results[i][j] = 0
			} else if slices.Equal(linesPrev[i-1], linesCurr[j-1]) {
				results[i][j] = 1 + results[i-1][j-1]
			} else {
				results[i][j] = max(results[i-1][j], results[i][j-1])
			}
		}
	}

	return results
}
