package lcs

import (
	"slices"

	"github.com/SergeyCherepiuk/share/client/pkg/internal"
)

func Diff(linesPrev, linesCurr []string) (deletions, insertions []int) {
	length := length(linesPrev, linesCurr)

	// TODO: Walk it from the beginning (and remove reverse at the end)
	i, j := len(linesPrev), len(linesCurr)
	for i != 0 && j != 0 {
		if linesPrev[i-1] == linesCurr[j-1] {
			i--
			j--
			continue
		}

		if length[i-1][j] <= length[i][j-1] {
			insertions = append(insertions, j-1)
			j--
		} else {
			deletions = append(deletions, i-1)
			i--
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

func length(linesPrev, linesCurr []string) internal.Matrix {
	results := make(internal.Matrix, len(linesPrev)+1)
	for i := range results {
		results[i] = make([]int, len(linesCurr)+1)
	}

	for i, result := range results {
		for j := range result {
			if i == 0 || j == 0 {
				results[i][j] = 0
			} else if linesPrev[i-1] == linesCurr[j-1] {
				results[i][j] = 1 + results[i-1][j-1]
			} else {
				results[i][j] = max(results[i-1][j], results[i][j-1])
			}
		}
	}

	return results
}
